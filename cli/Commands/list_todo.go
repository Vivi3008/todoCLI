package commands

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/alexeyco/simpletable"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func ListAllTodos(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "List all to do's",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := usecase.ListAll()

			if err != nil {
				fmt.Println(err)
			}

			if len(todos) == 0 {
				fmt.Println("You don't have to do's", chalk.ResetColor)
			} else {
				t := simpletable.New()
				t.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "Id"},
						{Align: simpletable.AlignCenter, Text: "Description"},
						{Align: simpletable.AlignCenter, Text: "Status"},
						{Align: simpletable.AlignCenter, Text: "Priority"},
						{Align: simpletable.AlignCenter, Text: "Date"},
					},
				}

				for _, td := range todos {
					r := []*simpletable.Cell{
						{Align: simpletable.AlignRight, Text: td.Id},
						{Text: td.Description},
						{Align: simpletable.AlignRight, Text: string(td.Status)},
						{Align: simpletable.AlignRight, Text: string(td.Priority)},
						{Align: simpletable.AlignRight, Text: string(td.CreatedAt.Format("02-01-2006"))},
					}

					t.Body.Cells = append(t.Body.Cells, r)
				}

				t.SetStyle(simpletable.StyleDefault)
				fmt.Println(t.String())
			}
		},
	}
	return addCmd
}

func ListById(usecase usecase.TodoUsecase) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list-one",
		Short: "List todo by Id",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(chalk.Red, "Please type to do ID to list")
			} else {
				id := args[0]
				todo, err := usecase.ListTodoId(id)

				if err != nil {
					fmt.Println(chalk.Red, err)
				}

				t := tabby.New()
				t.AddHeader("Id", "Description", "Status", "Priority", "Date")
				t.AddLine(todo.Id, todo.Description, todo.Status, todo.Priority, todo.CreatedAt.Format("02-01-2006"))
				t.Print()
			}
		},
	}
	return addCmd
}
