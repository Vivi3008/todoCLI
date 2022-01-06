package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

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
	t.WriteFile(td, "database.json")
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

func (t TodoStore) WriteFile(todo domain.Todo, source string) error {
	listTodos, err := t.ListAllTodos()

	if err != nil {
		return err
	}

	listTodos = append(listTodos, todo)
	jsonTodo, err := json.Marshal(listTodos)

	if err != nil {
		return err
	}

	ioutil.WriteFile(source, jsonTodo, os.ModeAppend)

	return nil
}
