package main

import (
	"os"

	"github.com/dpurbosakti/go-native/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup logger
	// middlewares.SetupLogger()
	if true {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	httpServer := handlers.NewHTTPServer(":8080")
	err := httpServer.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run http server")
	}
}
