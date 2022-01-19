package commom

import (
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/spf13/cobra"
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
