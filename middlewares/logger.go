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

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response recorder to capture the status code
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		// Serve HTTP with the response recorder
		next.ServeHTTP(recorder, r)

		end := time.Now()
		latency := end.Sub(start)

		logger := log.Info()

		// Get the status code from the response recorder
		status := recorder.status

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
	})
}

// statusRecorder captures the HTTP status code
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
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
