package task

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", h.GetAllTasks)
		r.Post("/", h.CreateTask)
		r.Get("/{id}", h.GetTaskByID)
		r.Put("/{id}", h.UpdateTask)
		r.Delete("/{id}", h.DeleteTask)
	})
}
