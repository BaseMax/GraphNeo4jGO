package DTO

type (
	Status      string
	UserRequest struct {
		Username string `json:"username,omitempty" validate:"required,gt=6,lowercase"`
		Name     string `json:"name,omitempty" validate:"required,gt=6"`
		Email    string `json:"email,omitempty" validate:"required,email"`
		Password string `json:"password,omitempty" validate:"required,gt=8"`
		Gender   uint8  `json:"gender,omitempty" validate:"required,oneof=1 2 3"`
	}

	UserResponse struct {
		Status Status `json:"status"`
		ID     uint   `json:"id,omitempty"`
		Token  string `json:"token,omitempty"`
		Data   any    `json:"data,omitempty"`
	}
)

const (
	StatusFound   Status = "Found"
	StatusCreated Status = "Created"
	StatusDeleted Status = "Deleted"
	StatusUpdated Status = "Updated"
	StatusError   Status = "Error"
)
