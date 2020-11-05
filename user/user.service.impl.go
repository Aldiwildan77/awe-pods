package users

import (
	"errors"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	UserRepository UserRepository
}

// NewUserService for entry point of user service
func NewUserService(userRepository *UserRepository) UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (service *userServiceImpl) Create(request CreateUserRequest) (response CreateUserResponse, _ error) {
	user := Users{
		ID:          request.ID,
		Name:        request.Name,
		Username:    request.Username,
		Email:       request.Email,
		Password:    request.Password,
		Age:         request.Age,
		Gender:      request.Gender,
		PhotoURL:    request.PhotoURL,
		Description: request.Description,
	}

	result, _ := service.UserRepository.Insert(user)

	response = CreateUserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Age:         result.Age,
		Gender:      result.Gender,
		PhotoURL:    result.PhotoURL,
		Description: result.Description,
	}

	return
}

func (service *userServiceImpl) Read(id uuid.UUID) (response ReadUserResponse, _ error) {
	result, e := service.UserRepository.Read(id)
	if e != nil {
		return response, e
	}

	response = ReadUserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Age:         result.Age,
		Gender:      result.Gender,
		PhotoURL:    result.PhotoURL,
		Description: result.Description,
	}

	return response, nil
}

func (service *userServiceImpl) Update(id uuid.UUID, request UpdateUserRequest) (response UpdateUserResponse, _ error) {
	_, e := service.UserRepository.Read(id)
	if e != nil {
		return response, e
	}

	user := Users{
		Name:        request.Name,
		Email:       request.Email,
		Gender:      request.Gender,
		PhotoURL:    request.PhotoURL,
		Description: request.Description,
	}

	result, e := service.UserRepository.Update(id, user)
	if e != nil {
		return response, e
	}

	response = UpdateUserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Age:         result.Age,
		Gender:      result.Gender,
		PhotoURL:    result.PhotoURL,
		Description: result.Description,
	}

	return response, nil
}

func (service *userServiceImpl) Delete(id uuid.UUID) (response DeleteUserResponse, _ error) {
	result, e := service.UserRepository.Read(id)
	if e != nil {
		return response, e
	}

	isDeleted := service.UserRepository.Delete(id)
	if isDeleted == false {
		return response, errors.New("failed to delete user")
	}

	response = DeleteUserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Username:    result.Username,
		Email:       result.Email,
		Age:         result.Age,
		Gender:      result.Gender,
		PhotoURL:    result.PhotoURL,
		Description: result.Description,
	}

	return response, nil
}

func (service *userServiceImpl) ValidateRegistered(validator UserValidator) bool {
	return service.UserRepository.ValidateUsernameOrEmail(validator.Username, validator.Email)
}
