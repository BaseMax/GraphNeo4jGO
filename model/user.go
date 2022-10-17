package model

type (
	Gender uint8
	User   struct {
		ID        uint
		Username  string
		Name      string
		Email     string
		Password  string
		Biography string
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
