package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/cli/commom"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func ListAllTodos(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list-all",
		Short: "List all to do's",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := usecase.ListAll()

			if err != nil {
				fmt.Println(err)
			}

			if len(todos) == 0 {
				fmt.Println(chalk.Yellow, "You don't have to do's", chalk.ResetColor)
			} else {
				commom.PrintTableTodos(todos)
			}
		},
	}
	return addCmd
}

func ListById(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "List todo by Id",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(chalk.Red, "Please type to do ID to list")
			} else {
				id := args[0]
				todo, err := usecase.ListTodoId(id)

				if err != nil {
					fmt.Println(chalk.Red, err, chalk.Reset)

				} else {
					commom.PrintTodo(todo)
				}
			}
		},
	}
	return addCmd
}
