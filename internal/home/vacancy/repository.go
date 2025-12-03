package vacancy

import (
	"context"
	"fmt"
	"time"

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
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, created_at) VALUES (@email, @role, @company, @salary, @type, @location, @created_at)`
	args := pgx.NamedArgs{
		"email":      vacancy.Email,
		"role":       vacancy.Role,
		"company":    vacancy.Company,
		"salary":     vacancy.Salary,
		"type":       vacancy.Type,
		"location":   vacancy.Location,
		"created_at": time.Now(),
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

func (r *VacancyRepository) GetVacancies() ([]*Vacancy, error) {
	query := `SELECT id, email, role, company, salary, type, location, created_at FROM vacancies;`
	rows, err := r.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancies from database: %w", err)
	}
	defer rows.Close()

	vacanсies, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Vacancy])
	if err != nil {
		return nil, fmt.Errorf("failed to collect vacancies from query: %w", err)
	}
	return vacanсies, nil
}