package product

import "pade-backend/pkg/entities"

type Service interface {
	GetProductByID(id string) (*entities.Product, error)
	GetAllProducts() ([]*entities.Product, error)
	GetProductsBySellerID(sellerID string) ([]*entities.Product, error)
	GetProductsByCategory(category string) ([]*entities.Product, error)
	CreateProduct(product *entities.Product) (*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(id string) error
	SearchProducts(query string) ([]*entities.Product, error)
}
