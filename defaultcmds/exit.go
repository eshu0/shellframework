package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)

//Exit the terminal
func Exit(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {
	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("exit() command called")
	return command.NewExitSuccessCommandResult()
}
