/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Vivi3008/todoCLI/cli/commom"
	"github.com/Vivi3008/todoCLI/domain"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/Vivi3008/todoCLI/store"
	"github.com/spf13/cobra"
)

var todoStore = store.NewTodoStore()
var todoUsecase = usecase.NewTodoUsecase(todoStore)

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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(Add())
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
