package cmd

import (
	"GraphNeo4jGO/config"
	"context"
	"time"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	repo, err := getRepo(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	service := getService(cfg, repo)
	_ = service
	return nil
}
