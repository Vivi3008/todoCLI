/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/Vivi3008/todoCLI/cli/cmd"
	"github.com/Vivi3008/todoCLI/domain/usecase"
	"github.com/Vivi3008/todoCLI/store"
)

var todoStore = store.NewTodoStore()
var todoUsecase = usecase.NewTodoUsecase(todoStore)

func main() {
	cmd.Execute(todoUsecase)
}
