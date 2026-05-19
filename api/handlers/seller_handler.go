package handlers

import (
	"pade-backend/api/presenter"
	"pade-backend/pkg/entities"
	"pade-backend/pkg/seller"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SellerHandler struct {
	sellerService seller.Service
}

func NewSellerHandler(sellerService seller.Service) *SellerHandler {
	return &SellerHandler{
		sellerService: sellerService,
	}
}

// CreateSeller godoc
// @Summary Create a new seller
// @Description Register as a seller
// @Accept json
// @Produce json
// @Param seller body entities.Seller true "Seller data"
// @Router /sellers [post]
func (h *SellerHandler) CreateSeller(c *fiber.Ctx) error {
	var seller entities.Seller
	if err := c.BodyParser(&seller); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.sellerService.CreateSeller(&seller)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusCreated, "Seller created successfully", result)
}

// GetSellerByID godoc
// @Summary Get seller by ID
// @Description Retrieve seller information
// @Produce json
// @Param id path string true "Seller ID"
// @Router /sellers/{id} [get]
func (h *SellerHandler) GetSellerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Seller ID is required")
	}

	seller, err := h.sellerService.GetSellerByID(id)
	if err != nil {
		return presenter.Error(c, fiber.StatusNotFound, "Seller not found")
	}

	return presenter.Success(c, fiber.StatusOK, "Seller retrieved successfully", seller)
}

// GetAllSellers godoc
// @Summary Get all sellers
// @Description Retrieve all sellers
// @Produce json
// @Router /sellers [get]
func (h *SellerHandler) GetAllSellers(c *fiber.Ctx) error {
	sellers, err := h.sellerService.GetAllSellers()
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Sellers retrieved successfully", sellers)
}

// GetNearestStores godoc
// @Summary Get nearest stores
// @Description Retrieve nearby stores based on coordinates
// @Produce json
// @Param lat query float64 true "Latitude"
// @Param lon query float64 true "Longitude"
// @Param limit query int false "Limit results"
// @Router /sellers/nearest [get]
func (h *SellerHandler) GetNearestStores(c *fiber.Ctx) error {
	lat := c.Query("lat")
	lon := c.Query("lon")
	limitStr := c.Query("limit", "10")

	if lat == "" || lon == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Latitude and longitude are required")
	}

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid latitude")
	}

	longitude, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid longitude")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	sellers, err := h.sellerService.GetNearestStores(latitude, longitude, limit)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Nearest stores retrieved successfully", sellers)
}

// UpdateSeller godoc
// @Summary Update seller
// @Description Update seller information
// @Accept json
// @Produce json
// @Param seller body entities.Seller true "Updated seller data"
// @Router /sellers/{id} [put]
func (h *SellerHandler) UpdateSeller(c *fiber.Ctx) error {
	var seller entities.Seller
	if err := c.BodyParser(&seller); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	seller.ID = c.Params("id")
	result, err := h.sellerService.UpdateSeller(&seller)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Seller updated successfully", result)
}

// DeleteSeller godoc
// @Summary Delete seller
// @Description Remove a seller
// @Param id path string true "Seller ID"
// @Router /sellers/{id} [delete]
func (h *SellerHandler) DeleteSeller(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Seller ID is required")
	}

	if err := h.sellerService.DeleteSeller(id); err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Seller deleted successfully", nil)
}
