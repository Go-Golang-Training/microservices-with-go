package service

import (
	"github.com/Go-Golang-Training/microservices-with-go/internal/domain"
	"github.com/Go-Golang-Training/microservices-with-go/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetProducts() []domain.Product {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductByID(id int) (*domain.Product, bool) {
	return s.repo.GetByID(id)
}

func (s *ProductService) AddProduct(p domain.Product) domain.Product {
	return s.repo.Create(p)
}
