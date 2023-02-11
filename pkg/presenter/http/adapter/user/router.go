package user

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/{id}", handler.Get)
		r.Post("/", handler.Create)
	})
	return r
}
