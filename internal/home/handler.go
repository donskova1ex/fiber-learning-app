package home

import (
	"bytes"
	"html/template"

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
		Bool("isAdmin", true).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("email", "a@a.ru").
		Msg("HomeHandler")
	tmpl := template.Must(template.ParseFiles("./html/page.html"))

	data := struct{ Count int }{Count: 1}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	return c.Send(tpl.Bytes())
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
