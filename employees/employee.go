package employees

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sql := `
	CREATE TABLE IF NOT EXISTS employees(
		id SERIAL PRIMARY KEY,
		full_name VARCHAR(255) NOT NULL,
		position VARCHAR(255) NOT NULL
	);
	`
	if _, err := conn.Exec(ctx, sql); err != nil {
		return err
	} else {
		return nil
	}
}

type Employee struct {
	ID       int
	FullName string
	Position string
}
