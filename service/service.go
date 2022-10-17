package service

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service/auth"
	"GraphNeo4jGO/service/user"
	"context"
	"github.com/go-playground/validator/v10"
)

type (
	Service interface {
		User() UserService
		Auth() Auth
		Tweet() Tweet
	}

	UserService interface {
		Login(ctx context.Context, user, pass string) (*DTO.UserResponse, error)
		Register(ctx context.Context, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Delete(ctx context.Context, id uint) (*DTO.UserResponse, error)
		Update(ctx context.Context, id uint, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Info(ctx context.Context, id string) (*DTO.UserResponse, error)
	}

	Auth interface {
		GenerateToken(id uint, username string) (string, error)
		ClaimsFromToken(token string) (any, error)
        BlackList(token string)
	}

	Tweet interface {
	}
)

type srv struct {
	user *user.ServiceImpl
	auth *auth.ServiceImpl
}

func (s *srv) User() UserService {
	return s.user
}

func (s *srv) Auth() Auth {
	return s.auth
}

func (s *srv) Tweet() Tweet {
	//TODO implement me
	panic("implement me")
}
func New(cfg *config.Config, repo repository.Repository) Service {
	authImpl := auth.New(cfg.Secrets, repo)
	userImpl := user.New(cfg, repo, validator.New(), authImpl)

	return &srv{
		user: userImpl,
		auth: authImpl,
	}
}
