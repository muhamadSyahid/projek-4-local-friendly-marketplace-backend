package seller

import "pade-backend/pkg/entities"

type Repository interface {
	Create(seller *entities.Seller) (*entities.Seller, error)
	GetByID(id string) (*entities.Seller, error)
	GetByUserID(userID string) (*entities.Seller, error)
	GetAll() ([]*entities.Seller, error)
	Update(seller *entities.Seller) (*entities.Seller, error)
	Delete(id string) error
	GetNearestStores(latitude, longitude float64, limit int) ([]*entities.Seller, error)
}
