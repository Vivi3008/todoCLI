package store

import (
	"errors"

	"github.com/Vivi3008/todoCLI/domain"
)

var (
	ErrIdNotExists = errors.New("id doesn't exists")
	ErrEmptyList   = errors.New("todo list is empty")
)

func (t TodoStore) ListAllTodos() []domain.Todo {
	var list []domain.Todo

	for _, todo := range t.todoStore {
		list = append(list, todo)
	}

	return list
}

func (t TodoStore) ListTodoById(id string) (domain.Todo, error) {
	listTodos := t.ListAllTodos()
	var td domain.Todo

	if len(listTodos) == 0 {
		return domain.Todo{}, ErrEmptyList
	}

	for _, todo := range listTodos {
		if todo.Id == id {
			td = todo
		}
	}

	if td.Id == "" {
		return domain.Todo{}, ErrIdNotExists
	}

	return td, nil
}
