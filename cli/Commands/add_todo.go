package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/cli/commom"
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/Vivi3008/todoCLI/store"
	"github.com/spf13/cobra"
)

var todoStore = store.NewTodoStore()
var todoUsecase = usecase.NewTodoUsecase(todoStore)

func Add() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to do",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Need todo description")
			}
			var description string
			for _, txt := range args {
				description += fmt.Sprintf("%s ", txt)
			}
			priority := commom.GetPriority(cmd)

			todo := domain.Todo{
				Description: description,
				Priority:    priority,
			}

			newTodo, err := todoUsecase.CreateTodo(todo)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(newTodo)
			}
		},
	}
	addCmd.Flags().BoolP("high", "H", false, "Define a todo normal priority.")
	addCmd.Flags().BoolP("low", "L", false, "Define a todo low priority.")
	return addCmd
}
