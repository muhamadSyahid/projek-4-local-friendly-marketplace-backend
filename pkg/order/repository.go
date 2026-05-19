package order

import "pade-backend/pkg/entities"

type Repository interface {
	Create(order *entities.Order) (*entities.Order, error)
	GetByID(id string) (*entities.Order, error)
	GetByBuyerID(buyerID string) ([]*entities.Order, error)
	GetBySellerID(sellerID string) ([]*entities.Order, error)
	Update(order *entities.Order) (*entities.Order, error)
	Delete(id string) error
	GetAll() ([]*entities.Order, error)
}
