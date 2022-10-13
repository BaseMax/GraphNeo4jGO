package cmd

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service"
	"context"
	"time"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repo, err := repository.New(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	srv := service.New(cfg, repo)
	_ = srv
	return nil
}
