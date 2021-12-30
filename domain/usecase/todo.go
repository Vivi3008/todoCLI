package usecase

import "github.com/Vivi3008/todoCLI/domain"

type TodoUsecase struct {
	tdUse domain.Repository
}

func NewTodoUsecase(td domain.Repository) TodoUsecase {
	return TodoUsecase{
		tdUse: td,
	}
}
