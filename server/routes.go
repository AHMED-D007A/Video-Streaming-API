package server

import (
	"Video-Streaming-API/services/authentication"
	"Video-Streaming-API/services/videoManagement"
	"Video-Streaming-API/services/videoStreaming"
	"Video-Streaming-API/services/videoUploading"
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
	uploadStorage := videoUploading.NewUploadStorage(db)
	uploadHandler := videoUploading.NewUploadHandler(uploadStorage)

	router.HandleFunc("/uvideos/upload", uploadHandler.UploadVideo).Methods("POST")
}

func RegisterVideoManagementRoutes(router *mux.Router, db *sql.DB) {
	videoStorage := videoManagement.NewVideoStorage(db)
	videoHandler := videoManagement.NewVideoHandler(videoStorage)

	router.HandleFunc("/uvideos", videoHandler.GetVideos).Methods("GET")
	router.HandleFunc("/uvideos/{id:[0-9]+}", videoHandler.UpdateVideo).Methods("PUT")
	router.HandleFunc("/uvideos/{id:[0-9]+}", videoHandler.DeleteVideo).Methods("DELETE")
}

func RegisterStreamingRoutes(router *mux.Router, db *sql.DB) {
	streamingStorage := videoStreaming.NewStreamingStorage(db)
	streamingHandler := videoStreaming.NewStreamingHandler(streamingStorage)

	router.HandleFunc("/stream/{videoID}/master.m3u8", streamingHandler.ServePlaylist).Methods("GET")
	router.PathPrefix("/stream/{videoID}/").HandlerFunc(streamingHandler.ServeSegment).Methods("GET")
}
