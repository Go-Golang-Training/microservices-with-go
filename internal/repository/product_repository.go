package repository

import "github.com/Go-Golang-Training/microservices-with-go/internal/domain"

type ProductRepository interface {
	GetAll() []domain.Product
	GetByID(id int) (*domain.Product, bool)
	Create(product domain.Product) domain.Product
}

type InMemoryProductRepository struct {
	products []domain.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []domain.Product{
			{ID: 1, Name: "Laptop", Price: 1499.99},
			{ID: 2, Name: "Mouse", Price: 49.99},
		},
	}
}

func (r *InMemoryProductRepository) GetAll() []domain.Product {
	return r.products
}

func (r *InMemoryProductRepository) GetByID(id int) (*domain.Product, bool) {
	for _, p := range r.products {
		if p.ID == id {
			return &p, true
		}
	}
	return nil, false
}

func (r *InMemoryProductRepository) Create(product domain.Product) domain.Product {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return product
}
