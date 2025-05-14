package task

import "errors"

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidID    = errors.New("invalid task ID")
	ErrTitleMissing = errors.New("title is required")
)
