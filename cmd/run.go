package cmd

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/controller/mux"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service"
	"context"
	"os"
	"os/signal"
	"syscall"
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

	rest := mux.New(srv, cfg)
	go func() {
		if err = rest.Start(); err != nil {
			panic(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	if err = rest.Stop(); err != nil {
		println("err: ", err)
	}

	return nil
}
