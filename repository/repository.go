package repository

import (
	"GraphNeo4jGO/config"
	"GraphNeo4jGO/model"
	"GraphNeo4jGO/repository/memcache"
	"GraphNeo4jGO/repository/neo4j"
	"GraphNeo4jGO/repository/postgres"
	"context"
)

type (
	Repository interface {
		UserPgx() PostgresUser
		UserGraph() GraphUser
		TweetGraph() GraphTweet
		Cache() Cache
	}

	PostgresUser interface {
		Create(ctx context.Context, u *model.User) (uint, error)
		User(ctx context.Context, id uint) (*model.User, error)
		UserFromUsername(ctx context.Context, username string) (*model.User, error)
		Update(ctx context.Context, u *model.User) error
		Delete(ctx context.Context, id uint) error
	}

	GraphUser interface {
		CreateUser(u model.GraphUser) error
		DeleteUser(user model.GraphUser) error
		UpdateUser(old, new string) (err error)
		UnFollowUser(u1 model.GraphUser, u2 model.GraphUser) (err error)
		FollowUser(u1 model.GraphUser, u2 model.GraphUser) (err error)
	}

	GraphTweet interface {
		NewTweet(t model.Tweet) ([16]byte, error)
	}

	Cache interface {
		Set(string, any)
		Get(string) (any, bool)
	}
)

type repo struct {
	user  *postgres.UserRepo
	graph *neo4j.Neo4j
	cache *memcache.CacheImpl
}

func (r *repo) UserPgx() PostgresUser {
	return r.user
}

func (r *repo) UserGraph() GraphUser {
	return r.graph
}

func (r *repo) TweetGraph() GraphTweet {
	return r.graph
}

func (r *repo) Cache() Cache {
	return r.cache
}

func New(ctx context.Context, cfg *config.Config) (Repository, error) {
	dbPool, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}
	u := postgres.NewUserRepo(dbPool)

	graph, err := neo4j.New(cfg.Neo4j)
	if err != nil {
		return nil, err
	}

	return &repo{
		user:  u,
		graph: graph,
		cache: memcache.New(),
	}, nil
}
