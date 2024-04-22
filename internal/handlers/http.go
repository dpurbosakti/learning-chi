package handlers

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/dpurbosakti/go-native/internal/handlers/account"
	"github.com/dpurbosakti/go-native/pkg/jsonhelper"
)

type HTTPServer struct {
	listenAddr string
}

func NewHTTPServer(listenaddr string) *HTTPServer {
	return &HTTPServer{
		listenAddr: listenaddr,
	}
}

func (s *HTTPServer) Run() error {
	router := http.NewServeMux()
	// _ = middlewares.ChainMiddleware(router, middlewares.LoggerMiddleware)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		jsonhelper.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "pong",
		})
	})
	accountHandler := account.NewAccountHandler()
	accountHandler.RegisterRoute(router)

	log.Info().Msgf("start http server at %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}
