package account

import (
	"net/http"

	"github.com/dpurbosakti/go-native/pkg/jsonhelper"
)

type AccountHandler struct {
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /accounts", h.Create)
	router.HandleFunc("GET /accounts/getall", h.GetAll)
	router.HandleFunc("GET /accounts/{id}", h.GetByID)
	router.HandleFunc("PUT /accounts/{id}", h.Update)
	router.HandleFunc("DELETE /accounts", h.Delete)
	// router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
	// 	jsonhelper.WriteJSON(w, http.StatusOK, map[string]any{
	// 		"message": "pong",
	// 	})
	// })
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	jsonhelper.WriteJSON(w, http.StatusOK, "ok")
}

func (h *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
