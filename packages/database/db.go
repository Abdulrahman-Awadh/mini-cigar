package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Database defines the contract for the database connection.
type Database interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Close(ctx context.Context) error
}

type CockroachDB struct {
	*pgx.Conn
}

func NewCockroachDB(connectionString string) (*CockroachDB, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return &CockroachDB{conn}, nil
}

func (db *CockroachDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.Conn.QueryRow(ctx, query, args...)
}

func (db *CockroachDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Conn.Exec(ctx, query, args...)
}

func (db *CockroachDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.Conn.Query(ctx, query, args...)
}

func (db *CockroachDB) Close(ctx context.Context) error {
	return db.Conn.Close(ctx)
}
