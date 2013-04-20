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

func NewApplication() *Application {
	app := new(Application)
	app.Add(NewHelpCommand())

	return app
}

func (self *Application) Add(cmd *Command) {
	if self.commands == nil {
		self.commands = make(map[string]*Command)
	}

	self.commands[cmd.GetName()] = cmd
}

func (self *Application) Run(output output.Output) (result bool) {
	self.output = output

	args, _ := self.parse()
	if len(args) == 0 {
		self.output.Error("Missing command")
		return false
	}

	self.command = args[0]
	return self.execute()
}

func (self *Application) GetOutput() output.Output {
	return self.output
}

func (self *Application) GetCommands() map[string]*Command {
	return self.commands
}

func (self *Application) execute() (result bool) {
	if cmd, ok := self.commands[self.command]; ok {
		self.output.Write("Executing command @{!}%s\n", output.NOTICE, self.command)

		cmd.Run(self)
	}

	return false
}

func (self *Application) parse() ([]string, error) {

	parser := flags.NewParser(&self.options, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)
	return parser.Parse()
}
