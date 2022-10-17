package mux

func (r *rest) routing() {
	api := r.router.PathPrefix("/api/v1").Subrouter()
	api.Use(r.handlers.loggerMiddleware)
	{
		user := api.PathPrefix("/user").Subrouter()
		user.Handle("/register/", handler(r.handlers.register))
		user.Handle("/login/", handler(r.handlers.login))
		user.Handle("/delete/", r.handlers.authorizationMiddleware(handler(r.handlers.delete)))
		user.Handle("/info/{username:[A-Za-z]+}", r.handlers.authorizationMiddleware(handler(r.handlers.myInfo)))
	}
	{
		tweets := api.PathPrefix("/tweets").Subrouter()
		tweets.Use(r.handlers.authorizationMiddlewareMux)
	}
}
