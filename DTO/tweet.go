package DTO

import "time"

type (
	TweetRequest struct {
		Username string `json:"username"`
		Text     string `json:"text" validate:"max=1024"`
	}

	TweetResponse struct {
		Data    any    `json:"data,omitempty"`
		Status  Status `json:"status"`
		TweetID string `json:"tweet_id,omitempty"`
	}

	CommentRequest struct {
		Text      string `json:"text"`
		Commenter string `json:"commenter"`
		Poster    string `json:"poster"`
		TweetID   string `json:"tweet_id"`
		UUID      string `json:"uuid"`
	}

	CommentResponse struct {
		Date      time.Time `json:"comment_time"`
		Data      any       `json:"data"`
		Status    Status    `json:"status"`
		TweetID   string    `json:"id"`
		Text      string    `json:"comment_text"`
		Poster    string    `json:"tweet_poster"`
		Commenter string    `json:"comment_poster"`
	}
)
