package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgreSQL(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
