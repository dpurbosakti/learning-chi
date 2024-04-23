package handlers

import (
	"net/http"

	store "github.com/dpurbosakti/go-native/internal/repositories"
	"github.com/dpurbosakti/go-native/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type HTTPServer struct {
	listenAddr string
}

func NewHTTPServer(listenaddr string) *HTTPServer {
	return &HTTPServer{
		listenAddr: listenaddr,
	}
}

func (server *HTTPServer) Run(s store.Store) error {
	router := chi.NewRouter()
	router.Use(middlewares.LoggerMiddleware)
	router.HandleFunc("/ping", Ping)
	h := SetRoutes(router, s)

	log.Info().Msgf("start http server at %s", server.listenAddr)
	return http.ListenAndServe(server.listenAddr, h)
}
