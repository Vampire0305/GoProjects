package task

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sudarshanmg/gotask/pkg/validation"
)

var validate = validator.New()

type TaskService interface {
	Create(req CreateTaskRequest) (*TaskResponse, error)
	GetAll(page, limit int, filter TaskFilter) ([]TaskResponse, int64, int, error)
	GetById(id int64) (*TaskResponse, error)
	Update(id int64, req UpdateTaskRequest) error
	Delete(id int64) error
}

type taskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func mapTasktoResponse(task *Task) TaskResponse {
	res := TaskResponse{
		ID:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return res
}

func (s *taskService) Create(req CreateTaskRequest) (*TaskResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, validation.FormatValidationError(err)
	}

	task := Task{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := s.repo.Create(&task)

	if err != nil {
		return nil, err
	}

	res := mapTasktoResponse(&task)
	return &res, nil
}

func (s *taskService) GetAll(page, limit int, filter TaskFilter) ([]TaskResponse, int64, int, error) {
	offset := (page - 1) * limit

	tasks, err := s.repo.FindAll(offset, limit, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	validateSortFields := map[string]bool{
		"id": true, "title": true, "created_at": true, "updated_at": true,
	}

	if !validateSortFields[filter.SortBy] {
		filter.SortBy = "id"
	}

	if filter.Order != "asc" && filter.Order != "desc" {
		filter.Order = "asc"
	}

	total, err := s.repo.CountAll()
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	responses := make([]TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		res := mapTasktoResponse(&t)
		responses = append(responses, res)
	}

	return responses, total, totalPages, nil
}

func (s *taskService) GetById(id int64) (*TaskResponse, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}

	task, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrNotFound
	}

	res := mapTasktoResponse(task)
	return &res, nil
}

func (s *taskService) Update(id int64, req UpdateTaskRequest) error {
	if id <= 0 {
		return ErrInvalidID
	}

	if err := validate.Struct(req); err != nil {
		return validation.FormatValidationError(err)
	}
	task, err := s.repo.FindById(id)

	if err != nil {
		return err
	}

	if task == nil {
		return ErrNotFound
	}
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	task.UpdatedAt = time.Now()
	return s.repo.Update(task)
}

func (s *taskService) Delete(id int64) error {
	if id <= 0 {
		return ErrInvalidID
	}
	task, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if task == nil {
		return ErrNotFound
	}
	return s.repo.Delete(id)

}
