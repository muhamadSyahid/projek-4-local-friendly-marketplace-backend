package order

import "pade-backend/pkg/entities"

type serviceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repo: repository,
	}
}

func (s *serviceImpl) CreateOrder(order *entities.Order) (*entities.Order, error) {
	return s.repo.Create(order)
}

func (s *serviceImpl) GetOrderByID(id string) (*entities.Order, error) {
	return s.repo.GetByID(id)
}

func (s *serviceImpl) GetOrdersByBuyerID(buyerID string) ([]*entities.Order, error) {
	return s.repo.GetByBuyerID(buyerID)
}

func (s *serviceImpl) GetOrdersBySellerID(sellerID string) ([]*entities.Order, error) {
	return s.repo.GetBySellerID(sellerID)
}

func (s *serviceImpl) UpdateOrderStatus(id string, status string) (*entities.Order, error) {
	order, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	order.Status = entities.ParseOrderStatus(status)
	return s.repo.Update(order)
}

func (s *serviceImpl) CancelOrder(id string) error {
	return s.repo.Delete(id)
}

func (s *serviceImpl) GetAllOrders() ([]*entities.Order, error) {
	return s.repo.GetAll()
}
