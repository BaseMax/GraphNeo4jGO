package tweet

import (
	"GraphNeo4jGO/repository"

	"github.com/go-playground/validator/v10"
)

type TweetService struct {
	repo     repository.Repository
	validate *validator.Validate
}

func New(repo repository.Repository, v *validator.Validate) *TweetService {
	return &TweetService{
		repo:     repo,
		validate: v,
	}
}
