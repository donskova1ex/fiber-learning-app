package vacancy

import (
	"go-fiber/pkg/t_adapter"
	"go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	handler := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}

	vacancyGroup := handler.router.Group("/vacancy")
	vacancyGroup.Post("/", handler.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")

	var component templ.Component
	if email == "" {
		component = components.Notification("Не задан e-mail", components.NotificationFail)
		return t_adapter.Render(c, component)
	}
	component = components.Notification("Vacancy created", components.NotificationSuccess)
	return t_adapter.Render(c, component)
}
