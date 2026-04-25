package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:docker@postgres-db:5432/postgres")

	return conn, err
}
