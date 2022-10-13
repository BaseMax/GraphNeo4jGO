package postgres

import (
	"GraphNeo4jGO/config"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Postgres) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(ctx, cfg.URI)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
