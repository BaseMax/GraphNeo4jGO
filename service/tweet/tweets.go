package tweet

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/model"
	"context"

	"github.com/gofrs/uuid"
)

func (t *TweetService) NewTweet(ctx context.Context, request DTO.TweetRequest) (*DTO.TweetResponse, error) {
	if err := t.validate.StructCtx(ctx, request); err != nil {
		return nil, err
	}

	tuidd, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	tweet := model.Tweet{
		Poster: request.Username,
		Text:   request.Text,
		UUID:   tuidd.String(),
	}

	id, err := t.repo.TweetGraph().NewTweet(tweet)
	if err != nil {
		return nil, err
	}

	return &DTO.TweetResponse{
		Status:  DTO.StatusCreated,
		TweetID: string(id[:]),
	}, nil
}

func (t *TweetService) UserTweets(username string, limit, skip int) (*DTO.TweetResponse, error) {
	tweets, err := t.repo.TweetGraph().UserTweets(username, limit, skip)
	if err != nil {
		return nil, err
	}

	return &DTO.TweetResponse{
		Status: DTO.StatusFound,
		Data:   tweets,
	}, nil
}

func (t *TweetService) UserTweet(username, uuid string) (DTO.TweetResponse, error) {
	if err := t.validate.Var(uuid, "uuid4"); err != nil {
		return DTO.TweetResponse{}, err
	}

	tweet, err := t.repo.TweetGraph().GetTweet(username, uuid)
	if err != nil {
		return DTO.TweetResponse{}, err
	}

	return DTO.TweetResponse{
		Status: DTO.StatusFound,
		Data:   tweet,
	}, nil
}

func (t *TweetService) DeleteTweet(username, uuid string) (DTO.TweetResponse, error) {
	if err := t.validate.Var(uuid, "uuid4"); err != nil {
		return DTO.TweetResponse{}, err
	}

	err := t.repo.TweetGraph().Delete(username, uuid)
	if err != nil {
		return DTO.TweetResponse{}, err
	}

	return DTO.TweetResponse{
		Status: DTO.StatusDeleted,
	}, nil
}
