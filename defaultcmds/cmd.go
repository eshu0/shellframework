package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)


// Command
func Cmd(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()
	log.LogPrintln("cmd() command called")
	env := shell.GetEnvironment()

	items, exists := env.GetVariable(sfinterfaces.LastCommands)
	if !exists {
		shell.Println("Environment has no last commands")
	} else {
		for _, cmd := range items.GetValues() {
			log.LogPrintln("cmd() command called")
			shell.Println(cmd)
		}
	}

	return command.NewSuccessCommandResult("")
}
