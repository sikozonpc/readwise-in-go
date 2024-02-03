package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	store := NewStore(s.db)
	mailer := NewSendGridMailer(Envs.SendGridAPIKey, Envs.SendGridFromEmail)

	service := NewService(store, mailer)
	service.RegisterRoutes(subrouter)

	log.Println("Starting API server on ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
