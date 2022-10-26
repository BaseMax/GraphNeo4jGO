package model

import "time"

type (
	// Tweet struct {
	// 	Username      string    `mapstructure:"username"`
	// 	Text          string    `mapstructure:"text"`
	// 	UUID          [16]byte  `mapstructure:"uuid"`
	// 	PublishDate   time.Time `mapstructure:"date"`
	// 	// TweetComments []TweetComment
	// }
	//
	// TweetComment struct {
	// 	Poster      string
	// 	Commenter   string
	// 	TweetID     [16]byte
	// 	Text        string
	// 	CommentUUID [16]byte
	// 	PublishDate time.Time
	// }

	Tweet struct {
		UUID      [16]byte  `mapstructure:"uuid"`
		Poster    string    `mapstructure:"username"`
		Text      string    `mapstructure:"text"`
		Date      time.Time `mapstructure:"date"`
		LikeCount int       `mapstructure:"like_count"`
		// Comments  []Comment ``
	}

	Comment struct {
		UUID      [16]byte ``
		TweetID   [16]byte
		Text      string ``
		Commenter string ``
		Poster    string
		Date      time.Time ``
	}
)
