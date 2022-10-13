package cmd

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service"
	"GraphNeo4jGO/service/auth"
	"GraphNeo4jGO/service/user"
	"github.com/go-playground/validator/v10"
)

type srv struct {
	user *user.ServiceImpl
	auth *auth.ServiceImpl
}

func (s *srv) User() service.UserService {
	return s.user
}

func (s *srv) Auth() service.Auth {
	return s.auth
}

func getService(cfg *config.Config, repo repository.Repository) service.Service {
	authImpl := auth.New(cfg.Secrets)
	userImpl := user.New(cfg, repo.UserRepo(), validator.New(), authImpl)

	return &srv{
		user: userImpl,
		auth: authImpl,
	}
}
