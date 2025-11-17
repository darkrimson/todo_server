package api

import (
	"go1f/pkg/api/handlers"

	"go1f/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func InitRoute() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {

		r.Post("/signin", handlers.SigninHandler)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth)

			r.Get("/task", handlers.GetTaskHandler)
			r.Post("/task", handlers.AddTaskHandler)
			r.Put("/task", handlers.UpdateTaskHandler)
			r.Delete("/task", handlers.DeleteTaskHandler)
			r.Get("/tasks", handlers.GetTasksHandler)
			r.Post("/task/done", handlers.DoneTaskHandler)
		})

		r.Get("/nextdate", handlers.NextDayHandler)
	})

	return r
}
