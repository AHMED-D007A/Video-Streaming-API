package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	authRouter := router.PathPrefix("/api/v1/").Subrouter()
	uploadingRouter := router.PathPrefix("/api/v1/").Subrouter()
	videoManagementRouter := router.PathPrefix("/api/v1/").Subrouter()
	streamingRouter := router.PathPrefix("/api/v1/").Subrouter()

	router.Use(LogMW)
	uploadingRouter.Use(AuthMW)
	videoManagementRouter.Use(AuthMW)

	RegisterAuthenticationRoutes(authRouter, s.db)
	RegisterUploadRoutes(uploadingRouter, s.db)
	RegisterVideoManagementRoutes(videoManagementRouter, s.db)
	RegisterStreamingRoutes(streamingRouter, s.db)

	// Serve the thumbnails from the "uploads" directory
	router.PathPrefix("/api/v1/").Handler(http.StripPrefix("/api/v1/", http.FileServer(http.Dir("./"))))

	log.Println("Starting server on:", s.addr)
	return http.ListenAndServe(s.addr, router)
}
