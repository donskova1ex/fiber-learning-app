package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}
	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
