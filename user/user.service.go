package users

import "github.com/google/uuid"

// UserService for store user behaviors
type UserService interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
	Read(id uuid.UUID) (ReadUserResponse, error)
	Update(id uuid.UUID, request UpdateUserRequest) (UpdateUserResponse, error)
	Delete(id uuid.UUID) (DeleteUserResponse, error)
	ValidateRegistered(validator UserValidator) bool
}
