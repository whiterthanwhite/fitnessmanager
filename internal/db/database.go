package db

import (
	"context"

	"github.com/jackc/pgx/v4"
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
