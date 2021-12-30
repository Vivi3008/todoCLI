package usecase

import (
	"errors"

	"github.com/Vivi3008/todoCLI/domain"
)

var (
	ErrIdNotExists = errors.New("this todo doesnt exists")
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

func (tc TodoUsecase) UpdateTodo(id string) (domain.Todo, error) {
	tod, err := tc.tdUse.ListTodoById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	//mark todo as done or pending
	if tod.Status == domain.Pend {
		tod.Status = domain.Done
	} else {
		tod.Status = domain.Pend
	}

	err = tc.tdUse.StoreTodo(tod)

	if err != nil {
		return domain.Todo{}, err
	}

	return tod, nil
}

func (tc TodoUsecase) DeleteTodo(id string) error {
	err := tc.tdUse.DeleteTodoId(id)
	if err != nil {
		return err
	}
	return nil
}
