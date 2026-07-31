package main

import (
	"taskmanager/view"
)

func main() {
	cli := view.NewCLI()
	cli.StartCommandLineInterface()
}
