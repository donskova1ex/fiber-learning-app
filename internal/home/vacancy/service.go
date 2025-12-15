package vacancy

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)


type VacancyService struct {
	repo *VacancyRepository
	log  *zerolog.Logger
}

func NewVacancyService(repo *VacancyRepository, log *zerolog.Logger) *VacancyService {
	return &VacancyService{
		repo: repo,
		log:  log,
	}
}

func (s *VacancyService) CreateVacancy(ctx context.Context, form VacancyCreateForm) error {
	vacancy := NewVacancyFromCreateForm(form)
	if err := s.repo.CreateVacancy(vacancy); err != nil {
		s.log.Error().Err(err).Str("operation", "CreateVacancy").Msg("failed to create vacancy in service")
		return fmt.Errorf("internal error: %w", err)
	}
	return nil
}

func (s *VacancyService) GetVacancies(limit, offset int) ([]*Vacancy, error) {
	vacanсies, err := s.repo.GetVacancies(limit, offset)
	if err != nil {
		s.log.Error().Err(err).Str("operation", "GetVacancies").Msg("failed to get vacancies in service")
		return nil, fmt.Errorf("internal error: %w", err)
	}
	return vacanсies, nil
}

func (s *VacancyService) CountAllVacancies() int {
	count := s.repo.CountAllVacancies()
	return count
}
