package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Go-Golang-Training/microservices-with-go/internal/service"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(s *service.HealthService) *HealthHandler {
	return &HealthHandler{service: s}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := h.service.GetHealthStatus()

	response := map[string]string{"status": status}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
