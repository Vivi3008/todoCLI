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

var source = "database.json"

func (t TodoStore) StoreTodo(td domain.Todo) error {
	if td.Id == "" {
		return ErrStore
	}
	t.todoStore[td.Id] = td
	err := t.WriteFileList(td)

	if err != nil {
		return err
	}
	return nil
}

func (t TodoStore) DeleteTodoId(id string) error {
	_, err := t.ListTodoById(id)

	if err != nil {
		return err
	}

	list, err := t.ListAllTodos()

	if err != nil {
		return err
	}

	var newList []domain.Todo

	for _, v := range list {
		if v.Id != id {
			newList = append(newList, v)
		}
	}

	jsonTodo, err := json.Marshal(newList)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(source, jsonTodo, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}

func (t TodoStore) DeleteAll() error {
	err := ioutil.WriteFile(source, []byte{}, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}

func (t TodoStore) WriteFileList(todo domain.Todo) error {
	listTodos, err := t.ListAllTodos()

	if err != nil {
		return err
	}

	listTodos = append(listTodos, todo)
	jsonTodo, err := json.Marshal(listTodos)

	if err != nil {
		return err
	}
	err = os.Chmod(source, 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(source, jsonTodo, os.ModeAppend)

	if err != nil {
		return err
	}

	return nil
}
