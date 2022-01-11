package commom

import (
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/spf13/cobra"
)

func GetPriority(cmd *cobra.Command) domain.Priority {
	high, _ := cmd.Flags().GetBool("high")
	low, _ := cmd.Flags().GetBool("low")
	if high {
		return domain.High
	}
	if low {
		return domain.Low
	}
	return domain.Normal
}
