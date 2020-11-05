package users

import "github.com/google/uuid"

// UserRepository for store user database behavior
type UserRepository interface {
	Insert(user Users) (Users, error)
	Read(id uuid.UUID) (Users, error)
	Update(id uuid.UUID, user Users) (Users, error)
	Delete(id uuid.UUID) bool
	ValidateUsernameOrEmail(username string, email string) bool
}
