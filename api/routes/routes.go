package routes

import (
	"pade-backend/api/handlers"
	"pade-backend/api/middleware"
	"pade-backend/pkg/marketplace"
	"pade-backend/pkg/order"
	"pade-backend/pkg/product"
	"pade-backend/pkg/seller"
	"pade-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all routes
func Setup(app *fiber.App, userService user.Service, productService product.Service, orderService order.Service, marketplaceService marketplace.Service, sellerService seller.Service) {
	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)
	marketplaceHandler := handlers.NewMarketplaceHandler(marketplaceService)
	sellerHandler := handlers.NewSellerHandler(sellerService)

	// Public routes
	public := app.Group("/api")

	// Auth routes
	auth := public.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	// Product routes (public)
	products := public.Group("/products")
	products.Get("", productHandler.GetAllProducts)
	products.Get("/:id", productHandler.GetProductByID)
	products.Get("/search", productHandler.SearchProducts)

	// Marketplace routes (public)
	marketplaces := public.Group("/marketplaces")
	marketplaces.Get("", marketplaceHandler.GetAllMarketplaces)
	marketplaces.Get("/:id", marketplaceHandler.GetMarketplaceByID)

	// Seller routes (public)
	sellers := public.Group("/sellers")
	sellers.Get("", sellerHandler.GetAllSellers)
	sellers.Get("/:id", sellerHandler.GetSellerByID)
	sellers.Get("/nearest", sellerHandler.GetNearestStores)

	// Protected routes
	protected := app.Group("/api")
	protected.Use(middleware.AuthMiddleware)

	// User routes (protected)
	users := protected.Group("/users")
	users.Get("", userHandler.GetAllUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)

	// Product routes (protected)
	protectedProducts := protected.Group("/products")
	protectedProducts.Post("", middleware.RoleMiddleware("seller", "admin"), productHandler.CreateProduct)
	protectedProducts.Put("/:id", middleware.RoleMiddleware("seller", "admin"), productHandler.UpdateProduct)
	protectedProducts.Delete("/:id", middleware.RoleMiddleware("seller", "admin"), productHandler.DeleteProduct)

	// Marketplace routes (protected)
	protectedMarketplaces := protected.Group("/marketplaces")
	protectedMarketplaces.Post("", middleware.RoleMiddleware("seller", "admin"), marketplaceHandler.CreateMarketplace)
	protectedMarketplaces.Put("/:id", middleware.RoleMiddleware("seller", "admin"), marketplaceHandler.UpdateMarketplace)
	protectedMarketplaces.Delete("/:id", middleware.RoleMiddleware("seller", "admin"), marketplaceHandler.DeleteMarketplace)

	// Seller routes (protected)
	protectedSellers := protected.Group("/sellers")
	protectedSellers.Post("", middleware.RoleMiddleware("buyer", "seller", "admin"), sellerHandler.CreateSeller)
	protectedSellers.Put("/:id", middleware.RoleMiddleware("seller", "admin"), sellerHandler.UpdateSeller)
	protectedSellers.Delete("/:id", middleware.RoleMiddleware("seller", "admin"), sellerHandler.DeleteSeller)

	// Order routes (protected)
	orders := protected.Group("/orders")
	orders.Post("", orderHandler.CreateOrder)
	orders.Get("/:id", orderHandler.GetOrderByID)
	orders.Get("/buyer", orderHandler.GetOrdersByBuyerID)
	orders.Put("/:id/status", orderHandler.UpdateOrderStatus)
	orders.Delete("/:id", orderHandler.CancelOrder)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
