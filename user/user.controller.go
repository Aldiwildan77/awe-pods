package users

import (
	response "awepods/handler"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

// UserController store controller behavior
type UserController struct {
	UserService UserService
}

// NewUserController to init user controller
func NewUserController(userService *UserService) UserController {
	return UserController{UserService: *userService}
}

// Route entry
func (controller *UserController) Route(route *mux.Router) {
	route.HandleFunc("/api/users", controller.Create).Methods("POST")
	route.HandleFunc("/api/users/{id}", controller.Read).Methods("GET")
	route.HandleFunc("/api/users/{id}", controller.Update).Methods("PUT")
	route.HandleFunc("/api/users/{id}", controller.Delete).Methods("DELETE")
}

// Create user data
func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	request.ID = uuid.New()

	if err := ValidateCreateUserRequest(request); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	validator := UserValidator{
		Username: request.Username,
		Email:    request.Email,
	}

	if ok := controller.UserService.ValidateRegistered(validator); ok {
		message := errors.New("user already exist")
		response.ResponseWithError(w, http.StatusBadRequest, message.Error())
		return
	}

	result, _ := controller.UserService.Create(request)
	response.ResponseWithJSON(w, http.StatusCreated, result)
}

// Read user data
func (controller *UserController) Read(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := controller.UserService.Read(id)
	if err != nil {
		response.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	response.ResponseWithJSON(w, http.StatusOK, result)
}

// Update user data
func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var request UpdateUserRequest

	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := controller.UserService.Update(id, request)
	if err != nil {
		response.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	response.ResponseWithJSON(w, http.StatusOK, result)
}

// Delete user data
func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := controller.UserService.Delete(id)
	if err != nil {
		response.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	response.ResponseWithJSON(w, http.StatusOK, result)
}
