package console

import (
	"github.com/jessevdk/go-flags"
	"github.com/mcuadros/console/output"
)

type Application struct {
	command  string
	commands map[string]*Command
	options  options
	output   output.Output
}

type options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Help    []bool `short:"h" long:"help" description:"Display this help message."`
}

func NewApplication(name string) *Application {
	return new(Application)
}

func (self *Application) Run(output output.Output) (result bool) {
	self.output = output

	args, _ := self.parse()
	if len(args) != 1 {
		self.output.Error("Missing command")
		return false
	}

	self.command = args[0]
	return self.execute()
}

func (self *Application) Add(cmd *Command) {
	if self.commands == nil {
		self.commands = make(map[string]*Command)
	}

	self.commands[cmd.GetName()] = cmd
}

func (self *Application) execute() (result bool) {
	if cmd, ok := self.commands[self.command]; ok {
		self.output.Write("Executing command @{!}%s\n", output.NOTICE, self.command)

		cmd.Run(self.output)
	}

	return false
}

func (self *Application) parse() ([]string, error) {
	return flags.NewParser(&self.options, flags.HelpFlag|flags.PassDoubleDash).Parse()
}
