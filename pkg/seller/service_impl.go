package seller

import "pade-backend/pkg/entities"

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) CreateSeller(seller *entities.Seller) (*entities.Seller, error) {
	return s.repo.Create(seller)
}

func (s *serviceImpl) GetSellerByID(id string) (*entities.Seller, error) {
	return s.repo.GetByID(id)
}

func (s *serviceImpl) GetSellerByUserID(userID string) (*entities.Seller, error) {
	return s.repo.GetByUserID(userID)
}

func (s *serviceImpl) GetAllSellers() ([]*entities.Seller, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) UpdateSeller(seller *entities.Seller) (*entities.Seller, error) {
	return s.repo.Update(seller)
}

func (s *serviceImpl) DeleteSeller(id string) error {
	return s.repo.Delete(id)
}

func (s *serviceImpl) GetNearestStores(latitude, longitude float64, limit int) ([]*entities.Seller, error) {
	return s.repo.GetNearestStores(latitude, longitude, limit)
}
