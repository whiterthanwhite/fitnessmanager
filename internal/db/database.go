package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Conn struct {
	conn *pgx.Conn
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

func (conn *Conn) Close(ctx context.Context) {
	conn.conn.Close(ctx)
}
