package repository

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/model"
	"GraphNeo4jGO/repository/postgres"
	"context"
)

type (
	Repository interface {
		UserRepo() User
	}

	User interface {
		Create(ctx context.Context, u *model.User) (uint, error)
		User(ctx context.Context, id uint) (*model.User, error)
		UserFromUsername(ctx context.Context, username string) (*model.User, error)
		Update(ctx context.Context, u *model.User) error
		Delete(ctx context.Context, id uint) error
	}
)

type r struct {
	user *postgres.UserRepo
}

func (r *r) UserRepo() User {
	return r.user
}

func New(ctx context.Context, cfg config.Postgres) (Repository, error) {
	dbPool, err := postgres.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	u := postgres.NewUserRepo(dbPool)
	return &r{
		user: u,
	}, nil
}
