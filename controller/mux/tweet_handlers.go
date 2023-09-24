package mux

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/service/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handlers) newTweet(w http.ResponseWriter, r *http.Request) error {
	var req DTO.TweetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	claim := r.Context().Value("claims")
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	req.Username = claims.Username

	res, err := h.srv.Tweet().NewTweet(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, res)
}

func (h *handlers) userTweets(w http.ResponseWriter, r *http.Request) error {
	var (
		q             = r.URL.Query()
		username      string
		limit, offset int = 20, 0
		err           error
	)
	username = mux.Vars(r)["username"]

	if q.Has("limit") {
		limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			return err
		}
	}

	if q.Has("offset") {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			return err
		}
	}

	res, err := h.srv.Tweet().UserTweets(username, limit, offset)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusFound, res)
}

func (h *handlers) userTweet(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	res, err := h.srv.Tweet().UserTweet(vars["username"], vars["uuid"])
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusFound, res)
}

func (h *handlers) deleteTweet(w http.ResponseWriter, r *http.Request) error {
	claim := r.Context().Value("claims")
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	res, err := h.srv.Tweet().DeleteTweet(claims.Username, mux.Vars(r)["uuid"])
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, res)
}

func (h *handlers) likeTweet(w http.ResponseWriter, r *http.Request) error {
	var (
		v        = mux.Vars(r)
		username = v["username"]
		uuid     = v["uuid"]
		claim    = r.Context().Value("claims")
	)
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	res, err := h.srv.Tweet().LikeTweet(claims.Username, username, uuid)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, res)
}

func (h *handlers) unlikeTweet(w http.ResponseWriter, r *http.Request) error {
	var (
		v        = mux.Vars(r)
		username = v["username"]
		uuid     = v["uuid"]
		claim    = r.Context().Value("claims")
	)
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	res, err := h.srv.Tweet().UnLike(claims.Username, username, uuid)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, res)
}

func (h *handlers) addComment(w http.ResponseWriter, r *http.Request) error {
	var (
		v        = mux.Vars(r)
		username = v["username"]
		uuid     = v["uuid"]
		claim    = r.Context().Value("claims")
	)
	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	var req DTO.CommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	req.Commenter = claims.Username
	req.Poster = username
	req.TweetID = uuid

	res, err := h.srv.Tweet().CommentOn(req)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, res)
}

func (h *handlers) allComments(w http.ResponseWriter, r *http.Request) error {
	var (
		v        = mux.Vars(r)
		username = v["username"]
		uuid     = v["uuid"]
	)
	res, err := h.srv.Tweet().AllComments(username, uuid)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusFound, res)
}

func (h *handlers) deleteComment(w http.ResponseWriter, r *http.Request) error {
	var (
		v        = mux.Vars(r)
		username = v["username"]
		uuid     = v["uuid"]
		claim    = r.Context().Value("claims")
	)

	var claims = &auth.JwtClaims{}
	if claims != nil {
		claims = claim.(*auth.JwtClaims)
	} else {
		return newError(http.StatusBadRequest, fmt.Errorf("cant get jwt claims from context"))
	}

	var req DTO.CommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	res, err := h.srv.Tweet().DeleteComment(DTO.CommentRequest{
		Commenter: claims.Username,
		Poster:    username,
		TweetID:   uuid,
		UUID:      req.UUID,
	})

	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, res)
}
