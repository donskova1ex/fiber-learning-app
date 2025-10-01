package home

import (
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
	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {

	users := []Users{
		{Id: 1, Name: "Alex"},
		{Id: 2, Name: "Don"},
		{Id: 3, Name: "Bob"},
	}
	names := []string{"Anna", "John", "Mike", "Bob", "Don", "Alex", "Kate", "Max", "Sara"}

	data := struct {
		Users []Users
		Names []string
	}{
		Users: users,
		Names: names,
	}

	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
