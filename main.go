package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mcuadros/go-composer/command"
)

func main() {
	parser := flags.NewNamedParser("test", flags.Default)
	parser.AddCommand("info", "Get info about a package", "A example of a command", new(command.Info))

	// Parse flags
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
