package users

import "github.com/google/uuid"

// UserRepository for store user database behavior
type UserRepository interface {
	Insert(user Users) (Users, error)
	Read(id uuid.UUID) (Users, error)
	ValidateUsernameOrEmail(username string, email string) bool
}
