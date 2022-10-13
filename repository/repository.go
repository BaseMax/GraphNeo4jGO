package repository

import (
	"GraphNeo4jGO/model"
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
