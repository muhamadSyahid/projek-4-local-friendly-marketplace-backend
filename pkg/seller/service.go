package seller

import "pade-backend/pkg/entities"

type Service interface {
	CreateSeller(seller *entities.Seller) (*entities.Seller, error)
	GetSellerByID(id string) (*entities.Seller, error)
	GetSellerByUserID(userID string) (*entities.Seller, error)
	GetAllSellers() ([]*entities.Seller, error)
	UpdateSeller(seller *entities.Seller) (*entities.Seller, error)
	DeleteSeller(id string) error
	GetNearestStores(latitude, longitude float64, limit int) ([]*entities.Seller, error)
}
