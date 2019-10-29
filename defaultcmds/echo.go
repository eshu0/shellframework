package defaultcmds

import (
	"github.com/eshu0/shellframework"
	"github.com/eshu0/shellframework/interfaces"
)

//Echo the terminal
func echo(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("echo() command called")
	ec := command.GetCommandInput()
	shell.Println(ec.GetInputWithOutCommand())
	log.LogPrintlnf("echo(): command name: %s", ec.GetCommandName())
	log.LogPrintlnf("echo(): raw input: %s", ec.GetRawInput())
	log.LogPrintlnf("echo(): printing: %s", ec.GetInputWithOutCommand())
	return shellframework.NewSuccessCommandResult("")
}
