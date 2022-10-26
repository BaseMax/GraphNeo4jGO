package user

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/model"
	"GraphNeo4jGO/repository/postgres"
	"context"
	"errors"
	"fmt"
)

var ErrUsernameExists = errors.New("username already exists. try another username")

func (s *ServiceImpl) Login(ctx context.Context, username, pass string) (*DTO.UserResponse, error) {
	user, err := s.repo.UserPgx().UserFromUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	var token string
	if comparePassword(user.Password, pass) {
		token, err = s.auth.GenerateToken(user.ID, username)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("invalid password")
	}

	return &DTO.UserResponse{
		Status: DTO.StatusFound,
		ID:     user.ID,
		Token:  token,
	}, nil
}

func (s *ServiceImpl) Register(ctx context.Context, request *DTO.UserRequest) (*DTO.UserResponse, error) {
	if err := s.validate.StructCtx(ctx, request); err != nil {
		return nil, err
	}
	u, err := s.repo.UserPgx().UserFromUsername(ctx, request.Username)
	if err != postgres.ErrNoRowFound {
		return nil, ErrUsernameExists
	}
	if u != nil {
		return nil, ErrUsernameExists
	}

	pass, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userModel := &model.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Password: pass,
		Gender:   model.Gender(request.Gender),
	}

	id, err := s.repo.UserPgx().Create(ctx, userModel)
	if err != nil {
		return nil, err
	}
	// create user in graph db
	if err = s.repo.UserGraph().CreateUser(model.GraphUser{Username: request.Username}); err != nil {
		return nil, err
	}

	token, err := s.auth.GenerateToken(id, request.Username)
	if err != nil {
		return nil, err
	}

	return &DTO.UserResponse{
		Status: DTO.StatusCreated,
		ID:     id,
		Token:  token,
		Data:   nil,
	}, nil
}

func (s *ServiceImpl) Delete(ctx context.Context, id uint, username string) (*DTO.UserResponse, error) {
	if err := s.repo.UserPgx().Delete(ctx, id); err != nil {
		return nil, err
	}

	if err := s.repo.UserGraph().DeleteUser(model.GraphUser{Username: username}); err != nil {
		return nil, err
	}
	return &DTO.UserResponse{
		Status: DTO.StatusDeleted,
		ID:     id,
	}, nil
}

func (s *ServiceImpl) Update(ctx context.Context, id uint, request *DTO.UserRequest) (*DTO.UserResponse, error) {
	if err := s.validate.StructCtx(ctx, request); err != nil {
		return nil, err
	}

	user, err := s.repo.UserPgx().UserFromUsername(ctx, request.Username)
	if err != postgres.ErrNoRowFound {
		return nil, err
	}
	if user != nil {
		return nil, ErrUsernameExists
	}

	pass, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userModel := &model.User{
		ID:       id,
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Password: pass,
		Gender:   model.Gender(request.Gender),
	}

	if err = s.repo.UserPgx().Update(ctx, userModel); err != nil {
		return nil, err
	}

	if user.Username != request.Username {
		// update username in graph
		if err = s.repo.UserGraph().UpdateUser(user.Username, request.Username); err != nil {
			return nil, err
		}
	}

	return &DTO.UserResponse{
		Status: DTO.StatusUpdated,
	}, nil
}

func (s *ServiceImpl) Info(ctx context.Context, username string) (*DTO.UserResponse, error) {
	user, err := s.repo.UserPgx().UserFromUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &DTO.UserResponse{Status: DTO.StatusFound, Data: user}, nil
}
