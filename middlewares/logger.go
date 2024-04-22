package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ChainMiddleware(handler http.Handler, middleware func(http.HandlerFunc) http.HandlerFunc) http.Handler {
	return middleware(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Execute the next handler
		next(w, r)

		end := time.Now()
		latency := end.Sub(start)

		logger := log.Info()
		status := w.(interface {
			Status() int
		}).Status()

		if status != http.StatusOK {
			logger = log.Error()
		}

		// Log request details using Zerolog
		logger.
			Str("method", r.Method).
			Str("uri", r.RequestURI).
			Int("status", status).
			Dur("latency", latency).
			Msg("request handled")
	}
}

func SetupLogger() {
	runLogFile, _ := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackMarshaler = func(err error) interface{} { return err }
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}
