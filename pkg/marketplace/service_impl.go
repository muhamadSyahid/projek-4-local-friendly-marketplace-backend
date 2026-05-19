package marketplace

import "pade-backend/pkg/entities"

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) CreateMarketplace(marketplace *entities.Marketplace) (*entities.Marketplace, error) {
	return s.repo.Create(marketplace)
}

func (s *serviceImpl) GetMarketplaceByID(id string) (*entities.Marketplace, error) {
	return s.repo.GetByID(id)
}

func (s *serviceImpl) GetAllMarketplaces() ([]*entities.Marketplace, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) GetMarketplacesByOwnerID(ownerID string) ([]*entities.Marketplace, error) {
	return s.repo.GetByOwnerID(ownerID)
}

func (s *serviceImpl) GetMarketplacesByCategory(category string) ([]*entities.Marketplace, error) {
	return s.repo.GetByCategory(category)
}

func (s *serviceImpl) UpdateMarketplace(marketplace *entities.Marketplace) (*entities.Marketplace, error) {
	return s.repo.Update(marketplace)
}

func (s *serviceImpl) DeleteMarketplace(id string) error {
	return s.repo.Delete(id)
}
