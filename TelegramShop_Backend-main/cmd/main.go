// @title TelegramShop Backend API
// @version 1.0
// @description This is a backend API for TelegramShop application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://194.187.122.144:5656/
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"telegramshop_backend/internal/repository/basket"
	"telegramshop_backend/internal/repository/categories"
	"telegramshop_backend/internal/repository/comment"
	"telegramshop_backend/internal/repository/favorites"
	"telegramshop_backend/internal/repository/firms"
	"telegramshop_backend/internal/repository/marks"
	"telegramshop_backend/internal/repository/orders"
	"telegramshop_backend/internal/repository/prices"
	"telegramshop_backend/internal/repository/products"
	"telegramshop_backend/internal/repository/users"

	avgMarksService "telegramshop_backend/internal/service/avg_marks"
	basketService "telegramshop_backend/internal/service/basket"
	categoriesService "telegramshop_backend/internal/service/categories"
	commentService "telegramshop_backend/internal/service/comment"
	favoritesService "telegramshop_backend/internal/service/favorites"
	firmsService "telegramshop_backend/internal/service/firms"
	marksService "telegramshop_backend/internal/service/marks"
	ordersService "telegramshop_backend/internal/service/orders"
	pricesService "telegramshop_backend/internal/service/prices"
	productsService "telegramshop_backend/internal/service/products"
	usersService "telegramshop_backend/internal/service/users"

	"telegramshop_backend/internal/handler"
	"telegramshop_backend/pkg/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "telegramshop_backend/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	db, err := postgres.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := users.NewRepository(db)
	basketRepo := basket.NewRepository(db)
	favoritesRepo := favorites.NewRepository(db)
	ordersRepo := orders.NewRepository(db)
	marksRepo := marks.NewRepository(db)
	avgmarksRepo := marks.NewRepository(db)
	productsRepo := products.NewRepository(db)
	pricesRepo := prices.NewRepository(db)
	categoriesRepo := categories.NewRepository(db)
	firmsRepo := firms.NewRepository(db)
	commentRepo := comment.NewRepository(db)

	userService := usersService.NewService(userRepo)
	basketService := basketService.NewService(basketRepo)
	favoritesService := favoritesService.NewService(favoritesRepo)
	ordersService := ordersService.NewService(ordersRepo)
	marksService := marksService.NewService(marksRepo)
	AvgMarksService := avgMarksService.NewService(avgmarksRepo)
	productsService := productsService.NewService(productsRepo)
	commentService := commentService.NewService(commentRepo)
	firmsService := firmsService.NewService(firmsRepo)
	categoriesService := categoriesService.NewService(categoriesRepo)
	pricesService := pricesService.NewService(pricesRepo)

	h := handler.NewHandler(userService, favoritesService, basketService, ordersService, firmsService, pricesService, categoriesService, productsService, marksService, AvgMarksService, commentService)

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// API routes
	h.InitRouter(app)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Fiber Listen error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")
	_ = app.Shutdown()
}
