package console

func NewHelpCommand() *Command {
	cmd := NewCommand("help")
	cmd.SetDescription("Displays help for a given command")
	cmd.SetHelp("The help command displays help for a given command")

	var options struct {
		Command string `short:"c" long:"command" description:"Command name" required:"true"`
	}
	cmd.SetOptions(&options)

	var handler = func(app *Application) {
		commands := app.GetCommands()
		commands[options.Command].PrintHelp(app.GetOutput())

		app.GetOutput().Write("\n@{y}Available commands:\n", 10)
		for name, command := range commands {
			app.GetOutput().Write("  @{g}%-20s@{!}\t%s\n", 10, name, command.GetDescription())
		}
	}
	cmd.SetHandler(handler)

	return cmd
}
