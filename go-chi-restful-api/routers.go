package main

import (
	"github.com/go-chi/chi"
)

func postsRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", List)

	r.Post("/", Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(addContext)
		r.Get("/", Get)
	})

	return r
}
