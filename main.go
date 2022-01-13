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
		Priority:    domain.High,
	}

	got, err := tdUsecase.CreateTodo(todo)

	/* 	err := tdUsecase.DeleteTodo("beb57eac-e6b9-4567-b708-29a82dcf3303")*/
	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Printf("Todo cadastrado com sucesso: %v\n", got)
	}

	/* fmt.Println("todo deletado com sucesso") */
}
