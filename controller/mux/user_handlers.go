package mux

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/service/auth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *handlers) register(w http.ResponseWriter, r *http.Request) error {
	var req DTO.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return newError(http.StatusBadRequest, err)
	}

	response, err := h.srv.User().Register(r.Context(), &req)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, response)
}

func (h *handlers) login(w http.ResponseWriter, r *http.Request) error {
	var req DTO.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return newError(http.StatusBadRequest, err)
	}

	response, err := h.srv.User().Login(r.Context(), req.Username, req.Password)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, response)
}

func (h *handlers) delete(w http.ResponseWriter, r *http.Request) error {
	claim := r.Context().Value("claims")
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	response, err := h.srv.User().Delete(r.Context(), claims.ID, claims.Username)
	if err != nil {
		return err
	}
	h.srv.Auth().BlackList(r.Context().Value("token").(string))

	return writeJson(w, http.StatusOK, response)
}

func (h *handlers) myInfo(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	response, err := h.srv.User().Info(r.Context(), username)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusFound, response)
}
