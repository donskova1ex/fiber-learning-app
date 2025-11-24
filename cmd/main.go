package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/home/vacancy"
	"go-fiber/pkg/database"
	"go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func main() {

	config.Init()

	logCfg := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logCfg)
	customLogger.Info().Msg("configurations initializing")

	app := NewApp(
		fiber.New(),
		customLogger,
		dbConfig,
	)
	customLogger.Info().Msg("new application created")

	if err := app.Listen(":3000"); err != nil {
		customLogger.Error().Err(err).Msg("error listetning server")
	}
	customLogger.Info().Msg("application started at port 3000")

}

func NewApp(
	app *fiber.App,
	logger *zerolog.Logger,
	dbConfig *config.DatabaseConfig,
) *fiber.App {

	dbpool := database.CreateDBPool(dbConfig, logger)
	defer dbpool.Close()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())
	app.Static("public", "./public")

	//Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, logger)

	//Handlers
	home.NewHandler(app, logger)
	vacancy.NewHandler(app, logger, vacancyRepo)

	return app
}
