package handlers

import (
	"net/http"

	ah "github.com/dpurbosakti/go-native/internal/handlers/account"
	store "github.com/dpurbosakti/go-native/internal/repositories"
	as "github.com/dpurbosakti/go-native/internal/services/account"
	rh "github.com/dpurbosakti/go-native/pkg/routerhelper"
	"github.com/go-chi/chi/v5"
)

func SetRoutes(router *chi.Mux, s store.Store) http.Handler {
	//account
	accountService := as.NewAccountService(s)
	accountHandler := ah.NewAccountHandler(accountService)
	rh.RegCrud(router, "/account", accountHandler)

	return router
}
