package main

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/pkg/logger"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {

	config.Init()

	cfg := config.NewDatabaseConfig()
	logCfg := config.NewLogConfig()

	customLogger := logger.NewLogger(logCfg)
	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(s string) string {
			switch s {
			case "":
				return ""
			default:
				return strings.ToUpper(s)
			}
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})
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
