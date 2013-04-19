package console

import (
	"github.com/jessevdk/go-flags"
	"github.com/mcuadros/console/output"
)

type Command struct {
	name   string
	values interface{}
	parser *flags.Parser
}

func NewCommand(name string) *Command {
	cmd := new(Command)
	cmd.SetName(name)

	return cmd
}

func (self *Command) Run(output output.Output) (result bool) {
	args, err := self.parse()
	if err != nil {
		output.Error("%s", err)
		return false
	}

	for index, value := range args {
		output.Info("\tType: %s, URL: %s\n", index, value)
	}

	output.Info("\tType: %s\n", self.values)

	return true
}

func (self *Command) SetName(name string) {
	self.name = name
}

func (self *Command) GetName() (name string) {
	return self.name
}

func (self *Command) SetOptions(options interface{}) {
	self.values = options
	self.parser = flags.NewParser(self.values, flags.HelpFlag|flags.PassDoubleDash)
}

func (self *Command) GetOptions() []*flags.Option {
	return self.parser.Groups[0].Options
}

func (self *Command) parse() ([]string, error) {
	return self.parser.Parse()
}
