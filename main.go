package main

import (
	"log"
	"os"
	"pade-backend/api/routes"
	"pade-backend/pkg/infra/database"
	infra_repo "pade-backend/pkg/infra/repository"
	"pade-backend/pkg/marketplace"
	"pade-backend/pkg/order"
	"pade-backend/pkg/product"
	"pade-backend/pkg/seller"
	"pade-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env in project root (if present)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found at project root, using system environment variables")
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081, http://127.0.0.1:8081, http://localhost:3000, http://127.0.0.1:3000, https://editor.swagger.io, https://muhamadSyahid.github.io",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Initialize MongoDB connection
	err := database.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	defer func() {
		if err := database.DisconnectMongoDB(); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()

	db := database.GetDatabase()

	// Initialize repositories
	userRepo := infra_repo.NewUserRepository(db)
	productRepo := infra_repo.NewProductRepository(db)
	orderRepo := infra_repo.NewOrderRepository(db)
	marketplaceRepo := infra_repo.NewMarketplaceRepository(db)
	sellerRepo := infra_repo.NewSellerRepository(db)

	// Initialize services
	userService := user.NewService(userRepo)
	productService := product.NewService(productRepo)
	orderService := order.NewService(orderRepo)
	marketplaceService := marketplace.NewService(marketplaceRepo)
	sellerService := seller.NewService(sellerRepo)

	// Setup routes
	routes.Setup(app, userService, productService, orderService, marketplaceService, sellerService)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
