package main

import (
	"net/http"

	"github.com/mahiro72/go-user-api/pkg/presenter/http/router"
	"github.com/mahiro72/go-user-api/pkg/registry"
)

func main() {

	repo, _ := registry.NewRepository()
	// defer _
	r := router.New(repo)

	if err := http.ListenAndServe(":8080", r); err != nil {
		// hoge
	}
}
