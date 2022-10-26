package cmd

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/controller/mux"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/repository/neo4j"
	"GraphNeo4jGO/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {
    // main context to init repository
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	repo, err := repository.New(ctx, cfg)
	if err != nil {
		return err
	}

    // make sure neo4j server is available
    if err := repo.UserGraph().(*neo4j.Neo4j).Ping(ctx); err != nil {
        return err
    }

	srv := service.New(cfg, repo)
	_ = srv

	rest := mux.New(srv, cfg)
	go func() {
		if err = rest.Start(); err != nil {
			if err != http.ErrServerClosed {
                panic(err)
            }
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Printf("stoping router...")
	if err = rest.Stop(); err != nil {
		return err
	}
	log.Printf("closing graph database...")
	if err = repo.UserGraph().(*neo4j.Neo4j).Close(); err != nil {
		return err
	}

	return nil
}
