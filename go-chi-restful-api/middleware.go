package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

func addContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
