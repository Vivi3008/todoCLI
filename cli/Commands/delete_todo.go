package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/spf13/cobra"
)

func Delete(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove to do",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Need todo ID!")
			}

			id := args[0]
			err := usecase.DeleteTodo(id)

			if err != nil {
				fmt.Println(err)
			}
		},
	}
	return addCmd
}

func DeleteAll(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete all to do's",
		Run: func(cmd *cobra.Command, args []string) {
			err := usecase.DeleteAllTodos()

			if err != nil {
				fmt.Println(err)
			}
		},
	}
	return addCmd
}
