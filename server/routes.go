package server

import (
	"Video-Streaming-Platform/services/authentication"
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterAuthenticationRoutes(router *mux.Router, db *sql.DB) {
	authStorage := authentication.NewAuthStorage(db)
	authHandler := authentication.NewAuthHandler(authStorage)

	router.HandleFunc("/signup", authHandler.Signup).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
}

func RegisterUploadRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/videos/upload", nil).Methods("POST")
}
