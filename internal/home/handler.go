package home

import (
	"go-fiber/pkg/t_adapter"
	"go-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type Users struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	handler := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	handler.router.Get("/", handler.home)
	handler.router.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return t_adapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
