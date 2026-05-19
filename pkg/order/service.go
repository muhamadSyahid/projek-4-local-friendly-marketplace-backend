package order

import "pade-backend/pkg/entities"

type Service interface {
	CreateOrder(order *entities.Order) (*entities.Order, error)
	GetOrderByID(id string) (*entities.Order, error)
	GetOrdersByBuyerID(buyerID string) ([]*entities.Order, error)
	GetOrdersBySellerID(sellerID string) ([]*entities.Order, error)
	UpdateOrderStatus(id string, status string) (*entities.Order, error)
	CancelOrder(id string) error
	GetAllOrders() ([]*entities.Order, error)
}
