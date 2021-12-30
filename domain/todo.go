package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string
type Priority string

const (
	Pend Status = "Pending"
	Done Status = "Done"
)

const (
	High   Priority = "High"
	Low    Priority = "Low"
	Normal Priority = "Normal"
)

var (
	ErrEmptyValues = errors.New("description, status or priority can't be empty")
)

type Todo struct {
	Id          string
	Description string
	Status      Status
	Priority    Priority
	CreatedAt   time.Time
}

func NewTodo(td Todo) (Todo, error) {
	if td.Description == "" || td.Status == "" || td.Priority == "" {
		return Todo{}, ErrEmptyValues
	}

	return Todo{
		Id:          uuid.New().String(),
		Description: td.Description,
		Status:      td.Status,
		Priority:    td.Priority,
		CreatedAt:   time.Now(),
	}, nil
}
