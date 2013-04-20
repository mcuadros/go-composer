package main

import (
	"composer/command"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	parser := flags.NewNamedParser("test", flags.Default)
	parser.AddCommand("info", "Get info aobut a package", "A example of a command", new(command.Info))

	// Parse flags
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
