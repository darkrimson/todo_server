package api

import (
	"go1f/pkg/api/handlers"

	"github.com/go-chi/chi/v5"
)

func InitRoute() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", handlers.NextDayHandler)
		r.Post("/task", handlers.AddTaskHandler)
	})

	return r
}
