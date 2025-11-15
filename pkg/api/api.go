package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/username/go-final-project/pkg/api/handlers"
)

func InitRoute() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", handlers.NextDayHandler)
		r.Post("/task", handlers.AddTaskHandler)
	})

	return r
}
