package DTO

type (
	Status      string
	UserRequest struct {
		Username string `json:"username,omitempty"`
		Name     string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		Gender   uint8  `json:"gender,omitempty"`
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
)
