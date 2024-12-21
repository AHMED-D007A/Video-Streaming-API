package server

import (
	"Video-Streaming-API/services/authentication"
	uploading "Video-Streaming-API/services/videoUploading"
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
	uploadStorage := uploading.NewUploadStorage(db)
	uploadHandler := uploading.NewUploadHandler(uploadStorage)

	router.HandleFunc("/videos/upload", uploadHandler.UploadVideo).Methods("POST")
}
