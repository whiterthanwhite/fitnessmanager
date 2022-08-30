package db

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/whiterthanwhite/fitnessmanager/internal/fitnessdata"
)

type Conn struct {
	conn *pgx.Conn
}

type CommandTag struct {
	RowsAffected int64
	Info         string
}

func Connect(ctx context.Context, dbAddr string) (*Conn, error) {
	conn := &Conn{}
	var err error
	conn.conn, err = pgx.Connect(ctx, dbAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (conn *Conn) InitTables(ctx context.Context) (*CommandTag, error) {
	ct, err := conn.conn.Exec(ctx, "create table workout ("+
		"entry_no integer,"+
		"workout_date date,"+
		"name varchar(100),"+
		"approach_no smallint,"+
		"repeat_no smallint,"+
		"description varchar(100),"+
		"PRIMARY KEY (entry_no));")
	if err != nil {
		return nil, err
	}
	return &CommandTag{ct.RowsAffected(), ct.String()}, nil
}

func (conn *Conn) DropTables(ctx context.Context) (*CommandTag, error) {
	ct, err := conn.conn.Exec(ctx, "drop table workout;")
	if err != nil {
		return nil, err
	}
	return &CommandTag{ct.RowsAffected(), ct.String()}, nil
}

func (conn *Conn) TableExist(ctx context.Context) (bool, error) {
	var exists bool
	err := conn.conn.QueryRow(ctx, "select exists ("+
		"select * from information_schema.tables "+
		"where table_name = 'workout');").Scan(&exists)
	return exists, err
}

func (conn *Conn) Close(ctx context.Context) {
	conn.conn.Close(ctx)
}

func (conn *Conn) InsertRecord(ctx context.Context, record *fitnessdata.Record) (*CommandTag, error) {
	entryNo, err := conn.getLastEntryNo(ctx)
	/*if err != nil {
		return nil, err
	}*/
	ct, err := conn.conn.Exec(ctx, "insert into workout values ($1, $2, $3, $4, $5, $6);", entryNo,
		record.Date, record.Name, record.Take, record.Repetitions, record.Description)
	if err != nil {
		return nil, err
	}
	return &CommandTag{ct.RowsAffected(), ct.String()}, nil
}

func (conn *Conn) getLastEntryNo(ctx context.Context) (int, error) {
	var entryNo int = 0
	if err := conn.conn.QueryRow(ctx, "select entry_no from workout order by entry_no desc limit 1;").
		Scan(&entryNo); err != nil {
		return entryNo, err
	}
	return entryNo + 1, nil
}
