package main

import (
	"context"
	"log"
	"os"

	"github.com/barzaevhalid/sotovik/internal/api/user"
	"github.com/barzaevhalid/sotovik/internal/configs"

	"github.com/barzaevhalid/sotovik/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := configs.LoadConfig()
	ctx := context.Background()

	pool, err := db.NewPool(ctx, *cfg)

	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer pool.Close()

	app := fiber.New(fiber.Config{
		AppName: "Sotovik API",
	})

	userRepo := user.NewUserRepository(pool)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	api := app.Group("/api")
	userHandler.RegisterRoutes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
