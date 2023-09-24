package user

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/model"
	"context"
	"fmt"
)

func (s *ServiceImpl) Follow(ctx context.Context, u1, u2 string) (*DTO.UserResponse, error) {
	if len(u1) < 6 && len(u2) < 6 {
		return nil, fmt.Errorf("invalid username")
	}

	if err := s.repo.UserGraph().FollowUser(model.GraphUser{Username: u1}, model.GraphUser{Username: u2}); err != nil {
		return nil, err
	}

	return &DTO.UserResponse{
		Status: DTO.StatusUpdated,
	}, nil
}

func (s *ServiceImpl) UnFollow(ctx context.Context, u1, u2 string) (*DTO.UserResponse, error) {
	if len(u1) < 6 && len(u2) < 6 {
		return nil, fmt.Errorf("invalid username")
	}

	if err := s.repo.UserGraph().UnFollowUser(model.GraphUser{Username: u1}, model.GraphUser{Username: u2}); err != nil {
		return nil, err
	}

	return &DTO.UserResponse{
		Status: DTO.StatusUpdated,
	}, nil

}

func (s *ServiceImpl) GetFollowers(ctx context.Context, u1 string) (*DTO.UserResponse, error) {
	if len(u1) < 6 {
		return nil, fmt.Errorf("invalid username")
	}

	usernames, err := s.repo.UserGraph().GetFollowers(u1)
	if err != nil {
		return nil, err
	}

	return &DTO.UserResponse{
		Status: DTO.StatusFound,
		Data:   map[string]any{"username": usernames, "count": len(usernames)},
	}, nil
}
