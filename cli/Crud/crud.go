package crud

import (
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/Vivi3008/todoCLI/domain/usecase"
)

func AddTodo(tdUse usecase.TodoUsecase, todo domain.Todo) (domain.Todo, error) {
	newTodo, err := tdUse.CreateTodo(todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return newTodo, nil
}
