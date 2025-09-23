package main

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Init()

	cfg := config.NewDatabaseConfig()
	fmt.Println(cfg)
	home.NewHandler(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
