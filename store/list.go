package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/Vivi3008/todoCLI/domain"
)

var (
	ErrIdNotExists = errors.New("id doesn't exists")
	ErrEmptyList   = errors.New("todo list is empty")
)

func (t TodoStore) ListAllTodos() ([]domain.Todo, error) {
	var list []domain.Todo

	file, err := ioutil.ReadFile("database.json")

	if err != nil {
		return nil, err
	}

	json.Unmarshal(file, &list)

	return list, nil
}

func (t TodoStore) ListTodoById(id string) (domain.Todo, error) {
	listTodos, err := t.ListAllTodos()

	if err != nil {
		return domain.Todo{}, err
	}
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
