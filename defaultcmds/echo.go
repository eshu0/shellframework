package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)

type EchoCommand struct {
}

func (command EchoCommand) Register(shell sfinterfaces.IShell) {
	shell.RegisterNewCommand("echo", "Echo text to terminal", Echo)
}

//Echo the terminal
func Echo(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("echo() command called")
	ec := command.GetCommandInput()
	shell.Println(ec.GetInputWithOutCommand())
	log.LogPrintlnf("echo(): command name: %s", ec.GetCommandName())
	log.LogPrintlnf("echo(): raw input: %s", ec.GetRawInput())
	log.LogPrintlnf("echo(): printing: %s", ec.GetInputWithOutCommand())
	return command.NewSuccessCommandResult("")
}
