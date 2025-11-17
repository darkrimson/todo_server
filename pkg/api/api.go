package api

import (
	"go1f/pkg/api/handlers"

	"github.com/go-chi/chi/v5"
)

func InitRoute() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", handlers.NextDayHandler)
		r.Get("/tasks", handlers.GetTasksHandler)
		r.Get("/task", handlers.GetTaskHandler)
		r.Post("/task", handlers.AddTaskHandler)
		r.Post("/task/done", handlers.DoneTaskHandler)
		r.Put("/task", handlers.UpdateTaskHandler)
		r.Delete("/task", handlers.DeleteTaskHandler)

	})

	return r
}
