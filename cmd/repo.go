package cmd

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/repository/postgres"
	"context"
)

type r struct {
	user *postgres.UserRepo
}

func (r *r) UserRepo() repository.User {
	return r.user
}

func getRepo(ctx context.Context, cfg config.Postgres) (repository.Repository, error) {
	dbPool, err := postgres.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	u := postgres.NewUserRepo(dbPool)
	return &r{
		user: u,
	}, nil
}
