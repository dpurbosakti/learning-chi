package main

import (
	"context"
	"os"

	"github.com/dpurbosakti/go-native/internal/handlers"
	store "github.com/dpurbosakti/go-native/internal/repositories"
	"github.com/dpurbosakti/go-native/middlewares"
	"github.com/dpurbosakti/go-native/pkg/config"
	"github.com/dpurbosakti/go-native/pkg/migrationhelper"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	// Setup logger
	middlewares.SetupLogger()
	if config.ENVIRONMENT == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	migrationhelper.RunDBMigration(config.MigrationURL, config.DBSource)

	s := store.NewStore(connPool)

	httpServer := handlers.NewHTTPServer(":8080")
	err = httpServer.Run(s)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run http server")
	}
}
