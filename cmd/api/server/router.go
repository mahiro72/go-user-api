package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mahiro72/go-user-api/cmd/api/adapter/user"
	"github.com/mahiro72/go-user-api/pkg/registry"
)

func NewRouter(repo *registry.Repository) http.Handler {
	r := chi.NewRouter()

	setCommonMiddlewares(r)

	userHandler := user.NewHandler(repo)

	r.Route("/users", func(r chi.Router) {
		r.Mount("/", user.NewRouter(userHandler))
	})
	return r
}
