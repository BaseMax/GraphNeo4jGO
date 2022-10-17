package mux

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/repository/postgres"
	"GraphNeo4jGO/service/user"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (*handlers) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s\t%s\t%s\n", time.Now().Format(time.RFC3339), r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

// access to normal users
func (h *handlers) authorizationMiddlewareMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken, err := getAuthToken(r)
		if err != nil {
			writeJson(w, http.StatusUnauthorized, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
			return
		}

		claims, err := h.srv.Auth().ClaimsFromToken(reqToken)
		if err != nil {
			writeJson(w, http.StatusUnauthorized, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
			return
		}

		if claims != nil {
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (h *handlers) authorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken, err := getAuthToken(r)
		if err != nil {
			//writeJson(w, http.StatusUnauthorized, dto.Error{Status: dto.StatusError, Error: err.Error()})
			return
		}
		claims, err := h.srv.Auth().ClaimsFromToken(reqToken)
		if err != nil {
			writeJson(w, http.StatusBadRequest, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
			return
		}

		if claims != nil {
			ctx := context.WithValue(r.Context(), "claims", claims)
			next(w, r.WithContext(ctx))
		}
	}
}

func handler(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		if err := f(w, request); err != nil {
			var e Error
			if errors.As(err, &e) {
				writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
			} else {
				if errors.Is(err, user.ErrUsernameExists) {
					writeJson(w, http.StatusBadRequest, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
				} else if errors.Is(err, postgres.ErrNoRowFound) {
					writeJson(w, http.StatusNotFound, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
				} else if errors.Is(err, postgres.ErrNoRowsAffected) {
					writeJson(w, http.StatusBadRequest, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
				} else {
					writeJson(w, http.StatusInternalServerError, DTO.Error{Status: DTO.StatusError, Err: err.Error()})
				}

			}
		}
	}
}

func getAuthToken(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
		return "", fmt.Errorf("auth header not in proper format")
	}

	reqToken = strings.TrimSpace(splitToken[1])
	return reqToken, nil
}
