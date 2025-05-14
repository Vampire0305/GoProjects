package task

import (
	"time"
)

type Task struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"max=500"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
	Completed   *bool   `json:"completed,omitempty"`
}

type TaskResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskFilter struct {
	Completed *bool
	SortBy    string
	Order     string
}
