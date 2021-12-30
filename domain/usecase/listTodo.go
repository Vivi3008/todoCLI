package usecase

import "github.com/Vivi3008/todoCLI/domain"

func (tc TodoUsecase) ListAll() []domain.Todo {
	list := tc.tdUse.ListAllTodos()

	return list
}

func (tc TodoUsecase) ListTodoId(id string) (domain.Todo, error) {
	todo, err := tc.tdUse.ListTodoById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}
