package main

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/Vivi3008/todoCLI/store"
)

func main() {
	todoStore := store.NewTodoStore()
	tdUsecase := usecase.NewTodoUsecase(todoStore)

	todo := domain.Todo{
		Description: "Fazer café sem açúcar",
		Status:      domain.Pend,
		Priority:    domain.High,
	}

	got, err := tdUsecase.CreateTodo(todo)

	if err != nil {
		fmt.Printf("Error %s", err)
	}

	fmt.Printf("Todo cadastrado com sucesso: %v\n", got)
}
