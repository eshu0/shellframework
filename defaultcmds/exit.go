package defaultcmds

import (
	"github.com/eshu0/shellframework"
	"github.com/eshu0/shellframework/interfaces"
)

//Exit the terminal
func exit(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {
	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("exit() command called")
	return shellframework.NewExitSuccessCommandResult()
}
