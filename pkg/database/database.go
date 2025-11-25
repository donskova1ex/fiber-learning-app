package database

import (
	"context"
	"go-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDBPool(config *config.DatabaseConfig, log *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		log.Error().Msg("Failed to create connection pool")
		panic(err)
	}
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Error().Msg("Failed to ping database")
		panic(err)
	}
	log.Info().Msg("Connection created")
	return dbpool

}
