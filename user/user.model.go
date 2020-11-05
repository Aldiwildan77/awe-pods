package users

import "github.com/google/uuid"

// CreateUserRequest request data for create user
type CreateUserRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
}

// CreateUserResponse response data for create user
type CreateUserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
}

// ReadUserResponse response data for read user
type ReadUserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
}

// UpdateUserRequest request data for update user
type UpdateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

// UpdateUserResponse response data for update user
type UpdateUserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
}

// DeleteUserResponse response data for delete user
type DeleteUserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
}
