package main

import (
	"context"
	"log"
	"os"

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

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server OK")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port)) // <- блокирует поток
}
