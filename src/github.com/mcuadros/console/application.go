package console

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/mcuadros/console/output"
)

type Application struct {
	command  string
	commands map[string]*Command
	options  struct {
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	}
}

func (self *Application) Run(output output.Output) (result bool) {
	self.commands["test"] = new(Command)

	output.Emergency("Test")
	args, _ := self.parse()
	if len(args) != 1 {
		fmt.Printf("Missing command")
		return false
	}

	self.command = args[0]
	return self.execute()
}

func (self *Application) Add(name string, command *Command) {
	if self.commands == nil {
		self.commands = make(map[string]*Command)
	}

	self.commands[name] = command
}

func (self *Application) execute() (result bool) {
	if cmd, ok := self.commands[self.command]; ok {
		fmt.Printf("Executing command: %s", self.command)
		cmd.Run()
	}

	return false
}

func (self *Application) parse() ([]string, error) {
	return flags.NewParser(&self.options, flags.HelpFlag|flags.PassDoubleDash).Parse()
}

type Command struct {
	name    string
	options struct {
		Some []bool `short:"s" long:"some" description:"Show verbose debug information" required:"true"`
	}
}

func (self *Command) parse() ([]string, error) {
	return flags.NewParser(&self.options, flags.HelpFlag|flags.PassDoubleDash).Parse()
}

func (self *Command) Run() (result bool) {
	args, err := self.parse()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false
	}

	for index, value := range args {
		fmt.Printf("\tType: %s, URL: %s\n", index, value)
	}

	return true
}
