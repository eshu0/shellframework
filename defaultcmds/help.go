package defaultcmds

import (
	"github.com/eshu0/shellframework"
	"github.com/eshu0/shellframework/interfaces"
)

// Help command
func help(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("help() command called")
	for _, obj := range shell.GetCommands() {
		shell.Println(obj.GetName())
		log.LogPrintln(obj.GetName())
	}
	return shellframework.NewSuccessCommandResult("")
}
