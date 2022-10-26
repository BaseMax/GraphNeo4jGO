package DTO

type (
	TweetRequest struct {
		Username string
		Text     string `json:"text" validate:"max=1024"`
	}

	TweetResponse struct {
		Status  Status `json:"status"`
		TweetID string `json:"tweet_id,omitempty"`
		Data    any    `json:"data,omitempty"`
	}
)
