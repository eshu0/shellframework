package dcmds

import (
	"github.com/eshu0/shellframework/interfaces"
)

type CmdCommand struct {
}

func (command CmdCommand) Register(shell sfinterfaces.IShell) {
	shell.RegisterNewCommand("cmd", "Commands command - jinx!", Cmd)
}

// Command
func Cmd(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()
	env := shell.GetEnvironment()

	items, exists := env.GetVariable(sfinterfaces.LastCommands)
	if !exists {
		shell.Println("Environment has no last commands")
	} else {
		for _, cmd := range items.GetValues() {
			log.LogDebug("cmd()", " command called")
			shell.Println(cmd)
		}
	}

	return command.NewSuccessCommandResult("")
}
