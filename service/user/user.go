package user

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service"
	"github.com/go-playground/validator/v10"
)

type ServiceImpl struct {
	user     repository.User
	auth     service.Auth
	validate *validator.Validate
	cfg      *config.Config
}

func New(cfg *config.Config, repo repository.User, v *validator.Validate, auth service.Auth) *ServiceImpl {
	return &ServiceImpl{
		user:     repo,
		auth:     auth,
		validate: v,
		cfg:      cfg,
	}
}
