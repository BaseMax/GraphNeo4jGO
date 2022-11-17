package model

import "time"

type (
	Tweet struct {
		Date      time.Time `mapstructure:"date"`
		Poster    string    `mapstructure:"username"`
		Text      string    `mapstructure:"text"`
		LikeCount int       `mapstructure:"like_count"`
		UUID      string    `mapstructure:"uuid"`
	}

	Comment struct {
		Date      time.Time ``
		Text      string    ``
		Commenter string    ``
		Poster    string
		UUID      string ``
		TweetID   string
	}
)
