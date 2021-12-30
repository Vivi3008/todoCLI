package usecase

import (
	"errors"

	"github.com/Vivi3008/todoCLI/domain"
)

var (
	ErrIdExists = errors.New("this todo")
)

func (t TodoUsecase) CreateTodo(todo domain.Todo) (domain.Todo, error) {
	newTodo, err := domain.NewTodo(todo)

	if err != nil {
		return domain.Todo{}, err
	}

	err = t.tdUse.StoreTodo(newTodo)

	if err != nil {
		return domain.Todo{}, err
	}

	return newTodo, nil
}
