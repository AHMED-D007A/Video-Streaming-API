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

	// subRouter := router.PathPrefix("api/v1/").Subrouter()

	router.Use(LogMW)

	log.Println("Starting server on:", s.addr[1:])
	return http.ListenAndServe(s.addr, router)
}
