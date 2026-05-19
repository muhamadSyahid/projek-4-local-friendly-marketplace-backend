package marketplace

import "pade-backend/pkg/entities"

type Service interface {
	CreateMarketplace(marketplace *entities.Marketplace) (*entities.Marketplace, error)
	GetMarketplaceByID(id string) (*entities.Marketplace, error)
	GetAllMarketplaces() ([]*entities.Marketplace, error)
	GetMarketplacesByOwnerID(ownerID string) ([]*entities.Marketplace, error)
	GetMarketplacesByCategory(category string) ([]*entities.Marketplace, error)
	UpdateMarketplace(marketplace *entities.Marketplace) (*entities.Marketplace, error)
	DeleteMarketplace(id string) error
}
