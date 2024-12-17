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
	subRouter := router.PathPrefix("/api/v1/").Subrouter()

	router.Use(LogMW)
	subRouter.Use(AuthMW)

	RegisterAuthenticationRoutes(authRouter, s.db)
	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Video Streaming Platform API"))
	})

	log.Println("Starting server on:", s.addr[1:])
	return http.ListenAndServe(s.addr, router)
}
