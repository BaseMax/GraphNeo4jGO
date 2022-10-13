package postgres

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Postgres) (repository.User, error) {
	db, err := pgxpool.Connect(ctx, cfg.URI)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
