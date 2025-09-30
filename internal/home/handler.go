package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	handler := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	h.customLogger.Info().
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("email", "a@a.ru").
		Msg("HomeHandler")
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{
		Count:   10,
		IsAdmin: true,
		CanUse:  true,
	}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
