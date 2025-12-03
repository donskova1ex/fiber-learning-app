package vacancy

import (
	"go-fiber/pkg/t_adapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	service      *VacancyService
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, service *VacancyService) {
	handler := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		service:      service,
	}

	vacancyGroup := handler.router.Group("/vacancy")
	vacancyGroup.Post("/", handler.createVacancy)
	vacancyGroup.Get("/", handler.getAll)
}

func (h *VacancyHandler) getAll(c *fiber.Ctx) error {
	vacancies, err := h.service.GetVacancies()
	if err != nil {
		h.customLogger.Error().Err(err).Msg("getting vacancies error")
	}
	return c.JSON(vacancies)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Salary:   c.FormValue("salary"),
		Type:     c.FormValue("type"),
		Location: c.FormValue("location"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Не задано расположение"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Не задана компания"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Не задана заработная плата"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Не задана должность"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Не задан тип компании"},
	)
	time.Sleep(2 * time.Second)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return t_adapter.Render(c, component, fiber.StatusBadRequest)
	}

	err := h.service.CreateVacancy(c.Context(), form)
	if err != nil {
		h.customLogger.Error().Err(err).Msg("creating vacancy error")
		component = components.Notification("Internal server error", components.NotificationFail)
		return t_adapter.Render(c, component, fiber.StatusInternalServerError)
	}
	component = components.Notification("Vacancy created", components.NotificationSuccess)
	return t_adapter.Render(c, component, fiber.StatusOK)
}
