package service

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/repository"
	"GraphNeo4jGO/service/auth"
	"GraphNeo4jGO/service/tweet"
	"GraphNeo4jGO/service/user"
	"context"

	"github.com/go-playground/validator/v10"
)

type (
	Service interface {
		User() UserService
		Auth() Auth
		Tweet() Tweet
	}

	UserService interface {
		UserGraph
		Login(ctx context.Context, user, pass string) (*DTO.UserResponse, error)
		Register(ctx context.Context, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Delete(ctx context.Context, id uint, username string) (*DTO.UserResponse, error)
		Update(ctx context.Context, id uint, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Info(ctx context.Context, id string) (*DTO.UserResponse, error)
		// Logout(ctx context.Context, token string) (*DTO.UserResponse, error)
		// RefreshToken(ctx context.Context, token string)
	}

	UserGraph interface {
		// Follow adds a FOLLOWING relationship from u1 to u2
		Follow(ctx context.Context, u1, u2 string) (*DTO.UserResponse, error)
		// UnFollow removes FOLLOWING relationship from u1 to u2
		UnFollow(ctx context.Context, u1, u2 string) (*DTO.UserResponse, error)
		// GetFollowers returns a list of usernames that follows u1 (as Data field in *DTO.UserResponse)
		GetFollowers(ctx context.Context, u1 string) (*DTO.UserResponse, error)
	}

	Auth interface {
		GenerateToken(id uint, username string) (string, error)
		ClaimsFromToken(token string) (any, error)
		BlackList(token string)
	}

	Tweet interface {
		NewTweet(ctx context.Context, request DTO.TweetRequest) (*DTO.TweetResponse, error)
		UserTweets(username string, limit, skip int) (*DTO.TweetResponse, error)
		UserTweet(username, uuid string) (DTO.TweetResponse, error)
		DeleteTweet(username, uuid string) (DTO.TweetResponse, error)
	}
)

type srv struct {
	user  *user.ServiceImpl
	auth  *auth.ServiceImpl
	tweet *tweet.TweetService
}

func (s *srv) User() UserService {
	return s.user
}

func (s *srv) Auth() Auth {
	return s.auth
}

func (s *srv) Tweet() Tweet {
	return s.tweet
}
func New(cfg *config.Config, repo repository.Repository) Service {
	v := validator.New()
	authImpl := auth.New(cfg.Secrets, repo)
	userImpl := user.New(cfg, repo, v, authImpl)
	tweetImpl := tweet.New(repo, v)

	return &srv{
		user:  userImpl,
		auth:  authImpl,
		tweet: tweetImpl,
	}
}
