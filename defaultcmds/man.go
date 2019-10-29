package dcmds

import (
	"strings"

	"github.com/eshu0/shellframework"
	"github.com/eshu0/shellframework/interfaces"
)

func Man(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {
	shell := command.GetShell()
	log := *shell.GetLog()

	log.LogPrintln("man() command called")

	ci := command.GetCommandInput()
	args := ci.GetArgs()

	log.LogPrintlnf("man(): Number of args: %d", len(args))

	if len(args) >= 1 {
		lowername := strings.ToLower(strings.TrimSpace(args[0]))
		for _, command := range shell.GetCommands() {
			if strings.ToLower(command.GetName()) == lowername {
				log.LogPrintlnf("man(): Command '%s' matched", command.GetName())

				shell.Printlnf("Command Name: %s", command.GetName())
				shell.Printlnf("Description: %s", command.GetDescription())
				shell.Println("Flags:")

				flgs := command.GetFlags()
				flgs.Parse()

				parsedflags := flgs.Parsedflags()

				if args != nil && len(args) > 0 && parsedflags != nil {
					/*
						for _, sflag := range parsedflags {
							// flag set per flag?
							// this needs to be reviewed....

							flgset := sflag.GetFlagSet()
							//flgset.PrintDefaults()
							if flgset != nil {
								shell.LogPrintln("man(): flagset not nil - Parsing flagset")

								for _, arg := range flgset.Args() {
									shell.Println(arg)
									shell.LogPrintlnf("man(): Argument: %s", arg)
								}
							} else {
								shell.LogPrintln("man(): flagset was nil ")
							}
						}
					*/

				} else {
					log.LogPrintln("man(): Not parsing flagset")
				}
			}
		}
	} else {
		shell.Println("What manual page do you want?")
	}

	return shellframework.NewSuccessCommandResult("")
}
