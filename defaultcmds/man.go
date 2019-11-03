package dcmds

import (
	"strings"

	"github.com/eshu0/shellframework/interfaces"
)

type ManCommand struct {
}

func (command ManCommand) Register(shell sfinterfaces.IShell) {
	shell.RegisterNewCommand("man", "Manual command similar to linux/unix systems", Man)
}

func Man(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {
	shell := command.GetShell()
	log := *shell.GetLog()

	ci := command.GetCommandInput()
	args := ci.GetArgs()

	log.LogDebugf("man()", "Number of args: %d", len(args))

	if len(args) >= 1 {
		lowername := strings.ToLower(strings.TrimSpace(args[0]))
		for _, command := range shell.GetCommands() {
			if strings.ToLower(command.GetName()) == lowername {
				log.LogDebugf("man()", "Command '%s' matched", command.GetName())

				shell.Printlnf("Command Name: %s", command.GetName())
				shell.Printlnf("Description: %s", command.GetDescription())
				flgs := command.GetCommandFlags()

				shell.Println("Flags:")
				flgs.PrintUsage()

			}
		}
	} else {
		shell.Println("What manual page do you want?")
	}

	return command.NewSuccessCommandResult("")
}
