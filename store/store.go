package store

import (
	"github.com/Vivi3008/todoCLI/domain"
)

type TodoStore struct {
	todoStore map[string]domain.Todo
}

func NewTodoStore() TodoStore {
	td := make(map[string]domain.Todo)
	return TodoStore{
		todoStore: td,
	}
}
