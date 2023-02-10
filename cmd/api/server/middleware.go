package main

import (
	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func setCommonMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}
