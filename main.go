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
		Description: "Fazer sexo",
		Priority:    domain.High,
	}

	got, err := tdUsecase.CreateTodo(todo)

	/* 	err := tdUsecase.DeleteAllTodos() */
	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Printf("Todo cadastrado com sucesso: %v\n", got)
	}

	/* fmt.Println("todo deletado com sucesso") */
}
