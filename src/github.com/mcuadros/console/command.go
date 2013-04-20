package console

import (
	"github.com/jessevdk/go-flags"
	"github.com/mcuadros/console/output"
)

type Command struct {
	name        string
	description string
	help        string
	values      interface{}
	parser      *flags.Parser
	handler     func(*Application)
}

func NewCommand(name string) *Command {
	cmd := new(Command)
	cmd.SetName(name)

	return cmd
}

func (self *Command) Run(app *Application) (result bool) {
	_, err := self.parse()
	if err != nil {
		app.GetOutput().Error("Error:%s\n", err)
		self.PrintHelp(app.GetOutput())
		return false
	}

	self.handler(app)
	return true
}

func (self *Command) SetName(name string) {
	self.name = name
}

func (self *Command) GetName() (name string) {
	return self.name
}

func (self *Command) SetDescription(description string) {
	self.description = description
}

func (self *Command) GetDescription() string {
	return self.description
}

func (self *Command) SetHelp(help string) {
	self.help = help
}

func (self *Command) GetHelp() string {
	return self.help
}

func (self *Command) SetHandler(handler func(app *Application)) {
	self.handler = handler
}

func (self *Command) GetHandler() func(app *Application) {
	return self.handler
}

func (self *Command) SetOptions(options interface{}) {
	self.values = options
	self.parser = flags.NewParser(self.values, flags.PassDoubleDash)
}

func (self *Command) GetOptions() []*flags.Option {
	return self.parser.Groups[0].Options
}

func (self *Command) GetValues() interface{} {
	return self.values
}

func (self *Command) PrintHelp(output output.Output) {
	output.Write("@{y}Usage:\n", 10)
	output.Write("  %s ", 10, self.name)
	for _, option := range self.GetOptions() {
		output.Write("[%s] ", 10, option)
	}
	output.Write("\n\n", 10)

	output.Write("@{y}Options:\n", 10)
	for _, option := range self.GetOptions() {
		output.Write("  @{g}%-20s@{!}\t%s\n", 10, option, option.Description)
	}

	output.Write("\n", 10)
	output.Write("@{y}Help:\n  @{!}%s\n", 10, self.help)
}

func (self *Command) parse() ([]string, error) {
	return self.parser.Parse()
}
