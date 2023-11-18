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
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "show all todo")

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
		errPrint(err)
		err = todos.Store(todoFile)
		errPrint(err)
	case *delete > 0:
		err := todos.Delete(*delete)
		errPrint(err)
		err = todos.Store(todoFile)
		errPrint(err)
	case *list:
		todos.Print()
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
