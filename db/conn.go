package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxIface interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func Connect() (*pgxpool.Pool, error) {

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("db URL is empty")
	}

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}

	fmt.Println("connected")
	return dbPool, nil
}
