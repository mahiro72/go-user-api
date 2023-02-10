package main

import (
	"net/http"

	"github.com/mahiro72/go-user-api/pkg/registry"
)

func main() {

	repo, _ := registry.NewRepository()
	// defer _
	r := NewRouter(repo)

	if err := http.ListenAndServe(":8080", r); err != nil {
		// hoge
	}
}
