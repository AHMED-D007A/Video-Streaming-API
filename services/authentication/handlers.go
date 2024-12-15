package authentication

import (
	"Video-Streaming-Platform/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
	var userpayload UserPayload
	err := json.NewDecoder(r.Body).Decode(&userpayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	userpayload.Password, err = bcrypt.GenerateFromPassword([]byte(userpayload.Password), 14)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := utils.CreateToken(userpayload.Name, userpayload.Email)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.storage.Register(&userpayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	response := fmt.Sprintf("{\n\t\"token\": \"%v\"\n}", token)

	w.Write([]byte(response))
}

// Login is a handler function that logs in a user
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userpayload UserPayload
	err := json.NewDecoder(r.Body).Decode(&userpayload)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.storage.Verify(&userpayload)
	if err != nil {
		log.Print(err.Error())
		if err.Error() == "not Found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("This Email is Not Registered"))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, userpayload.Password)
	if err != nil {
		w.Write([]byte("The email or the password are wrong."))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := utils.CreateToken(userpayload.Name, userpayload.Email)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"token\": \"%v\"}", token)

	w.Write([]byte(response))
}
