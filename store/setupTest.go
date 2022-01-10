package store

import (
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var todoStore = NewTodoStore()

func RunBeforeTests() {
	todos := ListTodosMock()

	for _, v := range todos {
		_ = todoStore.StoreTodo(v)
	}
}

func ListTodosMock() []domain.Todo {
	return []domain.Todo{
		{
			Id:          uuid.New().String(),
			Description: "Dormir",
			Status:      domain.Pend,
			Priority:    domain.Normal},
		{
			Id:          uuid.New().String(),
			Description: "Comer chocolate",
			Status:      domain.Done,
			Priority:    domain.High,
		},
		{
			Id:          uuid.New().String(),
			Description: "Jogar Game",
			Status:      domain.Pend,
			Priority:    domain.Normal,
		},
	}
}
