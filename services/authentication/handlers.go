package authentication

import (
	"net/http"
)

type AuthHandler struct {
	storage *AuthStorage
}

func NewAuthHandler(storage *AuthStorage) *AuthHandler {
	return &AuthHandler{
		storage: storage,
	}
}

// Register is a handler function that creates a new user in the database
func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
}

// Login is a handler function that logs in a user
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
}
