package handlers

import (
	"pade-backend/api/presenter"
	"pade-backend/pkg/entities"
	"pade-backend/pkg/order"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService order.Service
}

func NewOrderHandler(orderService order.Service) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with items
// @Accept json
// @Produce json
// @Param order body entities.Order true "Order data"
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.orderService.CreateOrder(&order)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusCreated, "Order created successfully", result)
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Retrieve order details
// @Produce json
// @Param id path string true "Order ID"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Order ID is required")
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		return presenter.Error(c, fiber.StatusNotFound, "Order not found")
	}

	return presenter.Success(c, fiber.StatusOK, "Order retrieved successfully", order)
}

// GetOrdersByBuyerID godoc
// @Summary Get orders by user
// @Description Retrieve orders for a specific user
// @Produce json
// @Param userId query string true "User ID"
// @Router /orders/buyer [get]
func (h *OrderHandler) GetOrdersByBuyerID(c *fiber.Ctx) error {
	userID := c.Query("userId")
	if userID == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "User ID is required")
	}

	orders, err := h.orderService.GetOrdersByBuyerID(userID)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Orders retrieved successfully", orders)
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param status query string true "New status"
// @Router /orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Query("status")

	if id == "" || status == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Order ID and status are required")
	}

	result, err := h.orderService.UpdateOrderStatus(id, status)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Order status updated successfully", result)
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel an existing order
// @Param id path string true "Order ID"
// @Router /orders/{id} [delete]
func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Order ID is required")
	}

	if err := h.orderService.CancelOrder(id); err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Order cancelled successfully", nil)
}
