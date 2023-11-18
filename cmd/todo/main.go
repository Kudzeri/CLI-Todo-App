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

	add := flag.Bool("add", true, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	//delete := flag.Bool("delete", false, "delete a todo")

	flag.Parse()

	todos := &todo.Todos{}

	err := todos.Load(todoFile)
	errPrint(err)

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		errPrint(err)
	case *complete > 0:
		err := todos.Complete(*complete)
		err = todos.Store(todoFile)
		errPrint(err)
	default:
		fmt.Fprintln(os.Stdout, ErrInvalidCommand)
		os.Exit(0)
	}
}

func errPrint(err error) {
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
}
