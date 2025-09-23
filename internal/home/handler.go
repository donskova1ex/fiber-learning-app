package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}
	handler.router.Get("/", handler.home)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
