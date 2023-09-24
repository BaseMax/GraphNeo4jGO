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

		tweets.Handle("/new/", handler(r.handlers.newTweet)).Methods(http.MethodPost)
		tweets.Handle("/{username}/", handler(r.handlers.userTweets)).Methods(http.MethodGet)
		tweets.Handle("/{username}/{uuid}/", handler(r.handlers.userTweet)).Methods(http.MethodGet)
		tweets.Handle("/delete/{uuid}/", handler(r.handlers.deleteTweet)).Methods(http.MethodDelete)

		tweets.Handle("/{username}/{uuid}/like/", handler(r.handlers.likeTweet)).Methods(http.MethodPatch)
		tweets.Handle("/{username}/{uuid}/unlike/", handler(r.handlers.unlikeTweet)).Methods(http.MethodPatch)

		tweets.Handle("/{username}/{uuid}/comment/", handler(r.handlers.addComment)).Methods(http.MethodPost)
		tweets.Handle("/{username}/{uuid}/comment/", handler(r.handlers.allComments)).Methods(http.MethodGet)
		tweets.Handle("/{username}/{uuid}/comment/", handler(r.handlers.deleteComment)).Methods(http.MethodDelete)
	}
}
