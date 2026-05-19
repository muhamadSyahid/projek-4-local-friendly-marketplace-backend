package handlers

import (
	"pade-backend/api/presenter"
	"pade-backend/pkg/entities"
	"pade-backend/pkg/product"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Add a new product to marketplace
// @Accept json
// @Produce json
// @Param product body entities.Product true "Product data"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.productService.CreateProduct(&product)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusCreated, "Product created successfully", result)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieve product information
// @Produce json
// @Param id path string true "Product ID"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Product ID is required")
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		return presenter.Error(c, fiber.StatusNotFound, "Product not found")
	}

	return presenter.Success(c, fiber.StatusOK, "Product retrieved successfully", product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all available products
// @Produce json
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Products retrieved successfully", products)
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products by query
// @Produce json
// @Param query query string true "Search query"
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Search query is required")
	}

	products, err := h.productService.SearchProducts(query)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Products found", products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product information
// @Accept json
// @Produce json
// @Param product body entities.Product true "Updated product data"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	product.ID = c.Params("id")
	result, err := h.productService.UpdateProduct(&product)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Product updated successfully", result)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Remove a product
// @Param id path string true "Product ID"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "Product ID is required")
	}

	if err := h.productService.DeleteProduct(id); err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Product deleted successfully", nil)
}
