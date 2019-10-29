package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)

type HelpCommand struct {
}

func (command HelpCommand) Register(shell sfinterfaces.IShell) {
	shell.RegisterNewCommand("help", "Help command", Help)
}

// Help command
func Help(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()
	for _, obj := range shell.GetCommands() {
		shell.Println(obj.GetName())
		log.LogDebug("help()", obj.GetName())
	}
	return command.NewSuccessCommandResult("")
}
