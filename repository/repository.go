package repository

import (
	"GraphNeo4jGO/model"
	"context"
)

type (
	User interface {
		Create(ctx context.Context, u *model.User) (uint, error)
		User(ctx context.Context, id uint) (*model.User, error)
		UserWithUsername(ctx context.Context, username string) (*model.User, error)
		Update(ctx context.Context, u *model.User) error
		Delete(ctx context.Context, id uint) error
	}
)
