package commom

import (
	"fmt"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/alexeyco/simpletable"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func GetPriority(cmd *cobra.Command) domain.Priority {
	high, _ := cmd.Flags().GetBool("high")
	low, _ := cmd.Flags().GetBool("low")

	var priority domain.Priority

	switch {
	case high:
		priority = domain.High
	case low:
		priority = domain.Low
	default:
		priority = domain.Normal
	}
	return priority
}

func PrintTodo(todo domain.Todo) {
	t := tabby.New()
	t.AddHeader("Id", "Description", "Status", "Priority", "Date")
	t.AddLine(todo.Id, todo.Description, todo.Status, todo.Priority, todo.CreatedAt.Format("02-01-2006"))
	t.Print()
}

func PrintTableTodos(todos []domain.Todo) {
	t := simpletable.New()
	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Id"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Date"},
			{Align: simpletable.AlignCenter, Text: "Status"},
		},
	}
	for i, j := 0, len(todos)-1; i < j; i, j = i+1, j-1 {
		todos[i], todos[j] = todos[j], todos[i]
	}

	for _, td := range todos {
		var status string
		if td.Status == domain.Pend {
			status = fmt.Sprint(chalk.Yellow, td.Status, chalk.Reset)
		} else {
			status = fmt.Sprint(chalk.Green, td.Status, chalk.Reset)
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: td.Id},
			{Text: td.Description},
			{Align: simpletable.AlignRight, Text: string(td.Priority)},
			{Align: simpletable.AlignRight, Text: string(td.CreatedAt.Format("02-01-2006"))},
			{Align: simpletable.AlignRight, Text: status},
		}

		t.Body.Cells = append(t.Body.Cells, r)
	}
	t.SetStyle(simpletable.StyleDefault)
	fmt.Println(t.String())
}
