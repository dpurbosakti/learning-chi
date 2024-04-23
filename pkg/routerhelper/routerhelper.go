package routerhelper

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CrudBase interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func RegCrud(r *chi.Mux, path string, c CrudBase) {
	r.Route(path, func(r chi.Router) {
		r.Post("/", c.Create)
		r.Get("/", c.GetList)
		r.Get("/{id}", c.GetDetail)
		r.Patch("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})
}
