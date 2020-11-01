package users

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UserValidator to store validator data
type UserValidator struct {
	Username string
	Email    string
}

// ValidateCreateUserRequest to validate create user request struct
func ValidateCreateUserRequest(request CreateUserRequest) (err error) {
	err = validation.ValidateStruct(&request,
		validation.Field(&request.Age, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Username, validation.Required, validation.Length(6, 25)),
		validation.Field(&request.Gender, validation.Required),
		validation.Field(&request.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&request.PhotoURL, is.URL),
	)

	// if err != nil {
	// 	panic(handler.ValidationError{
	// 		Message: err.Error(),
	// 	})
	// }

	return
}
