package mux

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/controller"
	"GraphNeo4jGO/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type (
	handlers struct {
		srv service.Service
		cfg config.Secrets
	}
	rest struct {
		handlers *handlers
		server   *http.Server
		router   *mux.Router
		cfg      config.Server
	}

	handlerFunc func(w http.ResponseWriter, r *http.Request) error
)

func (r *rest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.router.ServeHTTP(w, req)
}

func (r *rest) Start() error {
	log.Println("Starting server on:", r.cfg.Addr)
	r.routing()
	r.server = &http.Server{
		Addr:         r.cfg.Addr,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return r.server.ListenAndServe()

}

func (r *rest) Stop() error {
	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	return r.server.Shutdown(ctx)
}

func New(srv service.Service, cfg *config.Config) controller.Rest {
	return &rest{
		router: mux.NewRouter(),
		cfg:    cfg.Server,
		handlers: &handlers{
			srv: srv,
			cfg: cfg.Secrets,
		},
	}
}
