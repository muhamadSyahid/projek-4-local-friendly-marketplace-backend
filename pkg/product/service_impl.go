package product

import "pade-backend/pkg/entities"

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) GetProductByID(id string) (*entities.Product, error) {
	return s.repo.GetByID(id)
}

func (s *serviceImpl) GetAllProducts() ([]*entities.Product, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) GetProductsBySellerID(sellerID string) ([]*entities.Product, error) {
	return s.repo.GetBySellerID(sellerID)
}

func (s *serviceImpl) GetProductsByCategory(category string) ([]*entities.Product, error) {
	return s.repo.GetByCategory(category)
}

func (s *serviceImpl) CreateProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.Create(product)
}

func (s *serviceImpl) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.Update(product)
}

func (s *serviceImpl) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *serviceImpl) SearchProducts(query string) ([]*entities.Product, error) {
	return s.repo.Search(query)
}
