package posgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PostgresStorage struct {
	conn *pgx.Conn
}

func Open(connString string) (*PostgresStorage, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connString)

	return &PostgresStorage{conn: conn}, err
}

func (db *PostgresStorage) Close() error {
	return db.conn.Close(context.Background())
}
