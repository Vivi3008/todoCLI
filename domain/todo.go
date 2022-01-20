package domain

import (
	"errors"
	"math/rand"
	"time"
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
	ErrEmptyValues = errors.New("description, or priority can't be empty")
)

type Todo struct {
	Id          string
	Description string
	Status      Status
	Priority    Priority
	CreatedAt   time.Time
}

func NewTodo(td Todo) (Todo, error) {
	if td.Description == "" || td.Priority == "" {
		return Todo{}, ErrEmptyValues
	}

	if td.Status == "" {
		td.Status = Pend
	}

	return Todo{
		Id:          StringWithCharset(3),
		Description: td.Description,
		Status:      td.Status,
		Priority:    td.Priority,
		CreatedAt:   time.Now(),
	}, nil
}

const charset = "12345678910"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length)
}
