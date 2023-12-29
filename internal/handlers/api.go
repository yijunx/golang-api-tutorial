package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/yijunx/golang-api-tutorial/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware, applied to all endpoints
	r.Use(chimiddle.StripSlashes) // strips trailing slash

	r.Route(
		"/account",
		func(router chi.Router) {
			router.Use(middleware.Authorization)
			router.Get("/coins", GetCoinBalance)
		},
	)
}
