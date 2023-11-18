package main

import (
	"flag"
	"fmt"
	todo "github.com/kudzeri/todo-app"
	"os"
)

const (
	todoFile          = ".todos.json"
	ErrInvalidCommand = "INVALID COMMAND"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	//complete := flag.Bool("complete", false, "complete a todo")
	//delete := flag.Bool("delete", false, "delete a todo")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, ErrInvalidCommand)
		os.Exit(1)
	}
}
