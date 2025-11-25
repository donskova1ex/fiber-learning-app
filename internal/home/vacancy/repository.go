package vacancy

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *VacancyRepository) CreateVacancy(vacancy *Vacancy) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location) VALUES (@email, @role, @company, @salary, @type, @location)`
	args := pgx.NamedArgs{
		"email":    vacancy.Email,
		"role":     vacancy.Role,
		"company":  vacancy.Company,
		"salary":   vacancy.Salary,
		"type":     vacancy.Type,
		"location": vacancy.Location,
	}
	result, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("failed to create vacancy in database: %w", err)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("failed to create vacancy in database: no rows affected")
	}
	return nil
}
