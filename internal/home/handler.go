package home

import (
	"go-fiber/internal/home/vacancy"
	"go-fiber/pkg/t_adapter"
	"go-fiber/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	vacancyService  *vacancy.VacancyService
}

type Users struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, vacancyService *vacancy.VacancyService, customLogger *zerolog.Logger) {
	handler := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		vacancyService: vacancyService,
	}
	handler.router.Get("/", handler.home)
	handler.router.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	vacancies, err := h.vacancyService.GetVacancies(PAGE_ITEMS, (page - 1) * PAGE_ITEMS)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	component := views.Main(vacancies)
	return t_adapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Some error")
}
