package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/cli/commom"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func Update(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "update",
		Short: "Update todo to done or pending",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(chalk.Red, "Please type to do ID to list")
			} else {
				id := args[0]
				todo, err := usecase.UpdateTodo(id)

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
