package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Go-Golang-Training/microservices-with-go/internal/handler"
	"github.com/Go-Golang-Training/microservices-with-go/internal/service"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()

	// Initialize dependencies
	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)

	// Define routes
	mux.HandleFunc("/health", healthHandler.HealthCheck)

	return &Server{mux: mux}
}

func (s *Server) Start() {
	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.mux))
}
