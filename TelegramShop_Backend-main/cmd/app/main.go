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

	"telegramshop_backend/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	if err := a.Run(os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
