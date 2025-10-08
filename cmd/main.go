package main

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init()

	cfg := config.NewDatabaseConfig()
	logCfg := config.NewLogConfig()

	customLogger := logger.NewLogger(logCfg)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	fmt.Println(cfg, logCfg)

	home.NewHandler(app, customLogger)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
