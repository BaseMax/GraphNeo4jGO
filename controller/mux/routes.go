package mux

import "net/http"

func (r *rest) routing() {
	api := r.router.PathPrefix("/api/v1").Subrouter()
	api.Use(r.handlers.loggerMiddleware)
	{
		user := api.PathPrefix("/user").Subrouter()
		user.Handle("/register/", handler(r.handlers.register)).Methods(http.MethodPost)
		user.Handle("/login/", handler(r.handlers.login)).Methods(http.MethodPost)
		user.Handle("/delete/", r.handlers.authorizationMiddleware(handler(r.handlers.delete))).Methods(http.MethodDelete)
		user.Handle("/info/{username}/", r.handlers.authorizationMiddleware(handler(r.handlers.myInfo))).Methods(http.MethodGet)
	}
	{
		tweets := api.PathPrefix("/tweets").Subrouter()
		tweets.Use(r.handlers.authorizationMiddlewareMux)
	}
}
