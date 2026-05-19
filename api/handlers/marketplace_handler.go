package handlers

import (
	"pade-backend/api/presenter"
	"pade-backend/pkg/entities"
	"pade-backend/pkg/marketplace"

	"github.com/gofiber/fiber/v2"
)

type MarketplaceHandler struct {
	marketplaceService marketplace.Service
}

func NewMarketplaceHandler(marketplaceService marketplace.Service) *MarketplaceHandler {
	return &MarketplaceHandler{
		marketplaceService: marketplaceService,
	}
}

// CreateMarketplace godoc
// @Summary Create a new marketplace
// @Description Add a new marketplace
// @Accept json
// @Produce json
// @Param marketplace body entities.Marketplace true "Marketplace data"
// @Router /marketplaces [post]
func (h *MarketplaceHandler) CreateMarketplace(c *fiber.Ctx) error {
	var marketplace entities.Marketplace
	if err := c.BodyParser(&marketplace); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.marketplaceService.CreateMarketplace(&marketplace)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusCreated, "Marketplace created successfully", result)
}

// GetMarketplaceByID godoc
// @Summary Get marketplace by ID
// @Description Retrieve marketplace information
// @Produce json
// @Param id path string true "Marketplace ID"
// @Router /marketplaces/{id} [get]
func (h *MarketplaceHandler) GetMarketplaceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Marketplace ID is required")
	}

	marketplace, err := h.marketplaceService.GetMarketplaceByID(id)
	if err != nil {
		return presenter.Error(c, fiber.StatusNotFound, "Marketplace not found")
	}

	return presenter.Success(c, fiber.StatusOK, "Marketplace retrieved successfully", marketplace)
}

// GetAllMarketplaces godoc
// @Summary Get all marketplaces
// @Description Retrieve all available marketplaces
// @Produce json
// @Router /marketplaces [get]
func (h *MarketplaceHandler) GetAllMarketplaces(c *fiber.Ctx) error {
	marketplaces, err := h.marketplaceService.GetAllMarketplaces()
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Marketplaces retrieved successfully", marketplaces)
}

// UpdateMarketplace godoc
// @Summary Update marketplace
// @Description Update marketplace information
// @Accept json
// @Produce json
// @Param marketplace body entities.Marketplace true "Updated marketplace data"
// @Router /marketplaces/{id} [put]
func (h *MarketplaceHandler) UpdateMarketplace(c *fiber.Ctx) error {
	var marketplace entities.Marketplace
	if err := c.BodyParser(&marketplace); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	marketplace.ID = c.Params("id")
	result, err := h.marketplaceService.UpdateMarketplace(&marketplace)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Marketplace updated successfully", result)
}

// DeleteMarketplace godoc
// @Summary Delete marketplace
// @Description Remove a marketplace
// @Param id path string true "Marketplace ID"
// @Router /marketplaces/{id} [delete]
func (h *MarketplaceHandler) DeleteMarketplace(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Marketplace ID is required")
	}

	if err := h.marketplaceService.DeleteMarketplace(id); err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Marketplace deleted successfully", nil)
}
