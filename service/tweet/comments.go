package tweet

import (
	"GraphNeo4jGO/DTO"
	"GraphNeo4jGO/model"
	"time"

	"github.com/gofrs/uuid"
)

// AllComments implements service.Tweet
func (t *TweetService) AllComments(username string, uuid string) ([]DTO.CommentResponse, error) {
	if err := t.validate.Var(uuid, "uuid4"); err != nil {
		return nil, err
	}

	comments, err := t.repo.TweetGraph().GetComments(username, uuid)
	if err != nil {
		return nil, err
	}
	var res = make([]DTO.CommentResponse, len(comments))

	for i := 0; i < len(comments); i++ {
		c := comments[i]
		r := res[i]
		r.Date = c.Date
		r.Poster = c.Poster
		r.TweetID = c.TweetID
		r.Commenter = c.Commenter
	}

	return res, nil
}

// CommentOn implements service.Tweet
func (t *TweetService) CommentOn(r DTO.CommentRequest) (*DTO.CommentResponse, error) {
	if err := t.validate.Struct(r); err != nil {
		return nil, err
	}

	v4, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	u, err := t.repo.TweetGraph().CommentOn(model.Comment{
		Date:      time.Now(),
		Text:      r.Text,
		Commenter: r.Commenter,
		Poster:    r.Poster,
		UUID:      v4.String(),
		TweetID:   r.TweetID,
	})

	if err != nil {
		return nil, err
	}

	return &DTO.CommentResponse{
		Data:   map[string]any{"comment_uuid": u},
		Status: DTO.StatusCreated,
	}, nil
}

// DeleteComment implements service.Tweet
func (t *TweetService) DeleteComment(r DTO.CommentRequest) (*DTO.CommentResponse, error) {
	if err := t.validate.Struct(r); err != nil {
		return nil, err
	}

	if err := t.repo.TweetGraph().DeleteComment(model.Comment{
		Date:      time.Time{},
		Text:      "",
		Commenter: r.Commenter,
		Poster:    r.Poster,
		UUID:      r.UUID,
		TweetID:   r.TweetID,
	}); err != nil {
		return nil, err
	}

	return &DTO.CommentResponse{
		Status:  DTO.StatusDeleted,
		TweetID: r.TweetID,
	}, nil
}
