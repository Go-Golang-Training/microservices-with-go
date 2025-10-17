package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Go-Golang-Training/microservices-with-go/internal/handler"
	"github.com/Go-Golang-Training/microservices-with-go/internal/repository"
	"github.com/Go-Golang-Training/microservices-with-go/internal/service"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()

	// Initialize Health Dependencies
	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)

	// Initialize Product Dependencies
	productRepo := repository.NewInMemoryProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Routes
	mux.HandleFunc("/health", healthHandler.HealthCheck)
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productHandler.GetProductByID(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	return &Server{mux: mux}
}

func (s *Server) Start() {
	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.mux))
}
