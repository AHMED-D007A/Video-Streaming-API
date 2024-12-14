package server

import (
	"Video-Streaming-Platform/services/authentication"
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterAuthenticationRoutes(router *mux.Router, db *sql.DB) {
	authStorage := authentication.NewAuthStorage(db)
	authHandler := authentication.NewAuthHandler(authStorage)

	router.HandleFunc("/register", authHandler.Signup).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
}
