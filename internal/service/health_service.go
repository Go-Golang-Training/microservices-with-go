package service

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) GetHealthStatus() string {
	// business logic could include DB checks, external APIs, etc.
	return "ok"
}
