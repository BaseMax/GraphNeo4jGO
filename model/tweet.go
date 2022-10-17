package model

type (
	Tweet struct {
		Username string
		Text     string
		UUID     [16]byte
	}
)
