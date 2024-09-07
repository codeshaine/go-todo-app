package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/codeshaine/go-todo-app"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as complete")
	incomplete := flag.Int("incomplete", 0, "mark a todo as incomplete")
	del := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()
	todos := todo.Todos{}

	//loading the json data into the struct list
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		save(&todos)
	case *complete > 0:
		err := todos.Compelete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		save(&todos)
	case *incomplete > 0:
		err := todos.Incompelete(*incomplete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		save(&todos)
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		save(&todos)

	case *list:
		todos.Print()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	scnaner := bufio.NewScanner(r)
	scnaner.Scan()
	if err := scnaner.Err(); err != nil {
		return "", err
	}
	if len(scnaner.Text()) == 0 {
		return "", fmt.Errorf("empty todo is not allowed")
	}
	return scnaner.Text(), nil
}

func save(todos *todo.Todos) {
	err := todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
