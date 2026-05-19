package product

import "pade-backend/pkg/entities"

type Repository interface {
	Create(product *entities.Product) (*entities.Product, error)
	GetByID(id string) (*entities.Product, error)
	GetAll() ([]*entities.Product, error)
	GetBySellerID(sellerID string) ([]*entities.Product, error)
	GetByCategory(category string) ([]*entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(id string) error
	Search(query string) ([]*entities.Product, error)
}
