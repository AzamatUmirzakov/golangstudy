package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
