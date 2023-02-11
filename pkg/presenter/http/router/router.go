package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mahiro72/go-user-api/pkg/presenter/http/adapter/user"
	"github.com/mahiro72/go-user-api/pkg/presenter/http/middleware"
	"github.com/mahiro72/go-user-api/pkg/registry"
)

func New(repo *registry.Repository) http.Handler {
	r := chi.NewRouter()

	middleware.SetCommonMiddlewares(r)

	userHandler := user.NewHandler(repo)

	r.Route("/users", func(r chi.Router) {
		r.Mount("/", user.NewRouter(userHandler))
	})
	return r
}
