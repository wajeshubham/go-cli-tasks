package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	task "github.com/wajeshubham/go-cli-tasks"
)

const (
	taskFile = "tasks.json"
)

func main() {
	add := flag.Bool("add", false, "Add new task")
	position := flag.Int("done", 0, "Mark task done")
	delete := flag.Int("delete", 0, "Delete a task")
	list := flag.Bool("list", false, "List all tasks")

	flag.Parse()
	tasks := &task.Tasks{}

	if err := tasks.Load(taskFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		tasks.Add(task)
		err = tasks.Store(taskFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		tasks.Print()
	case *position > 0:
		err := tasks.Complete(*position)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = tasks.Store(taskFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *delete > 0:
		err := tasks.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = tasks.Store(taskFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty task is not allowed")
	}

	return text, nil
}
