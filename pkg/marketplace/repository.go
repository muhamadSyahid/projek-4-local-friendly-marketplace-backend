package marketplace

import "pade-backend/pkg/entities"

type Repository interface {
	Create(marketplace *entities.Marketplace) (*entities.Marketplace, error)
	GetByID(id string) (*entities.Marketplace, error)
	GetAll() ([]*entities.Marketplace, error)
	GetByOwnerID(ownerID string) ([]*entities.Marketplace, error)
	GetByCategory(category string) ([]*entities.Marketplace, error)
	Update(marketplace *entities.Marketplace) (*entities.Marketplace, error)
	Delete(id string) error
}
