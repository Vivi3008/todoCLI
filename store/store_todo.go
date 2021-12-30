package store

import (
	"errors"

	"github.com/Vivi3008/todoCLI/domain"
)

var (
	ErrStore = errors.New("err: todo has missing values")
)

func (t TodoStore) StoreTodo(td domain.Todo) error {
	if td.Id == "" {
		return ErrStore
	}
	t.todoStore[td.Id] = td
	return nil
}

func (t TodoStore) DeleteTodoId(id string) error {
	var del bool
	for _, v := range t.todoStore {
		if v.Id == id {
			delete(t.todoStore, v.Id)
			del = true
		}
	}
	if !del {
		return ErrIdNotExists
	}
	return nil
}
