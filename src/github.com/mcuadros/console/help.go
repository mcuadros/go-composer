package console

func NewHelpCommand() *Command {
	cmd := NewCommand("help")

	var opts struct {
		Command string `short:"c" long:"command" description:"Command name" required:"true"`
	}

	var HelpHandler = func(app *Application) {
		commands := app.GetCommands()
		commands[opts.Command].PrintHelp(app.GetOutput())

		app.GetOutput().Write("\n@{y}Available commands:\n", 10)
		for name, command := range commands {
			app.GetOutput().Write("  @{g}%-20s@{!}\t%s\n", 10, name, command.GetDescription())
		}
	}

	cmd.SetDescription("Displays help for a given command")
	cmd.SetHelp("The help command displays help for a given command")
	cmd.SetHandler(HelpHandler)
	cmd.SetOptions(&opts)
	return cmd
}
