package main

import (
	"composer/command"
	"composer/misc"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	misc.GetOutput().Info("test")

	parser := flags.NewNamedParser("test", flags.Default)
	parser.AddCommand("info", "Get info about a package", "A example of a command", new(command.Info))

	// Parse flags
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
