package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)

type ExitCommand struct {
}

func (command ExitCommand) Register(shell sfinterfaces.IShell) {
	shell.RegisterNewCommand("exit", "Exit terminal command", Exit)
}

//Exit the terminal
func Exit(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {
	return command.NewExitSuccessCommandResult()
}
