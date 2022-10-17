package user

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service/auth"
	"github.com/go-playground/validator/v10"
)

type ServiceImpl struct {
	postgresUser repository.PostgresUser
	graphUser    repository.GraphUser
	auth         *auth.ServiceImpl
	validate     *validator.Validate
	cfg          *config.Config
}

func New(cfg *config.Config, repo repository.Repository, v *validator.Validate, auth *auth.ServiceImpl) *ServiceImpl {
	return &ServiceImpl{
		postgresUser: repo.UserPgx(),
		graphUser:    repo.UserGraph(),
		auth:         auth,
		validate:     v,
		cfg:          cfg,
	}
}
