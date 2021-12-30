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
