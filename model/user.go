package model

type (
	Gender uint8
	User   struct {
		ID       uint
		Username string
		Name     string
		Email    string
		Password string
		Gender   Gender
	}
)

const (
	_ Gender = iota
	Male
	Female
	Other
)
