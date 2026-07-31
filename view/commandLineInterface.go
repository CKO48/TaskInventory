package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"taskmanager/control"
)

type CLI struct {
	tc control.TaskController
}

var taskController *control.TaskController

func NewCLI() *CLI {
	return &CLI{tc: *control.NewTaskController()}
}

func (cli *CLI) StartCommandLineInterface() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to TaskInventory!")

	for {
		fmt.Print(">")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimRight(input, "\r\n")

		args, err := Parse(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		exitCode, _ := cli.handleCommand(args)
		if exitCode == 0 {
			break
		}
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <\"title\"> [-d <description>] - Add a new task with the specified title and optional description.")
	fmt.Println("  complete <\"title\"> - Mark a task as completed.")
	fmt.Println("  list - List all tasks or only completed tasks.")
	fmt.Println("  remove <\"title\"> - Remove a task with the specified title.")
	fmt.Println("  help - Display this help message.")
}

func (cli *CLI) handleCommand(args []string) (int, error) {
	if len(args) == 0 {
		fmt.Println("No command provided. Type 'help' for available commands.")
		return 1, nil
	}

	switch args[0] {
	case "add":
		message, err := cli.tc.AddTask(args)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

	case "list":
		message, err := cli.tc.ListTasks(args)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

	case "remove":
		message, err := cli.tc.RemoveTask(args)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

	case "complete":
		message, err := cli.tc.CompleteTask(args)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

	case "help":
		printHelp()

	case "exit":
		fmt.Println("Exiting TaskInventory. Goodbye!")
		return 0, nil

	default:
		fmt.Println("Unknown command. Type 'help' for available commands.")
	}

	return 1, nil
}
