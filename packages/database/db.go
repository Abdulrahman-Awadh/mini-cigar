package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToCockroachDB(connectionString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Database defines the contract for the database connection.
type Database interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

// CockroachDB represents the CockroachDB connection.
type CockroachDB struct {
	*pgxpool.Pool
}

// NewCockroachDB initializes a new CockroachDB connection pool.
func NewCockroachDB(connectionString string) (*CockroachDB, error) {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return &CockroachDB{pool}, nil

}
func (db *CockroachDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.Pool.QueryRow(ctx, query, args...)
}

// Exec executes a query and returns the command tag.
func (db *CockroachDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Pool.Exec(ctx, query, args...)
}
func (db *CockroachDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.Pool.Query(ctx, query, args...)
}
