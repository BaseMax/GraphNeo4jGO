package model

type (
	Gender uint8
	User   struct {
		Username  string
		Name      string
		Email     string
		Password  string
		Biography string
		ID        uint
		Gender    Gender
	}

	GraphUser struct {
		Username string
	}
)

const (
	_ Gender = iota
	Male
	Female
	Other
)
