package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/cli/commom"
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/spf13/cobra"
)

func Add(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to do",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Need todo description")
			}
			var description string
			for i, txt := range args {
				//remover espaço da ultima palavra
				if i == (len(args) - 1) {
					description += txt
					break
				}
				description += fmt.Sprintf("%s ", txt)
			}

			priority := commom.GetPriority(cmd)

			todo := domain.Todo{
				Description: description,
				Priority:    priority,
			}

			newTodo, err := usecase.CreateTodo(todo)

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
