package service

import (
	"GraphNeo4jGO/DTO"
	"context"
)

type (
	Service interface {
		User() UserService
		Auth() Auth
	}

	UserService interface {
		Login(ctx context.Context, user, pass string) (*DTO.UserResponse, error)
		Register(ctx context.Context, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Delete(ctx context.Context, id uint) (*DTO.UserResponse, error)
		Update(ctx context.Context, id uint, request *DTO.UserRequest) (*DTO.UserResponse, error)
		Info(ctx context.Context, id uint) (*DTO.UserResponse, error)
	}

	Auth interface {
		GenerateToken(id uint, username string) (string, error)
		ClaimsFromToken(token string) (any, error)
	}
)
