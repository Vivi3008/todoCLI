/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/Vivi3008/todoCLI/domain/usecase"
	commands "github.com/Vivi3008/todoCLI/task/Commands"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A cli to add and view pending to-do's",
	Long: `This program allow to add pending tasks to do, view list to-dos, mark as done and
	and delete to-dos`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(usecase usecase.TodoUsecase) {
	rootCmd.AddCommand(
		commands.Add(usecase),
		commands.Delete(usecase),
		commands.DeleteAll(usecase),
		commands.ListAllTodos(usecase),
		commands.ListById(usecase),
		commands.Update(usecase),
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", true, "Help message for toggle")
}
