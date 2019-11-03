package dcmds

import (
	"strings"

	"github.com/eshu0/shellframework/interfaces"
)

type EnvCommand struct {
}

func (command EnvCommand) Register(shell sfinterfaces.IShell) {

	//fg := command.GetFlags()
	//Flags := []sfinterfaces.IFlag{}

	//flags := []sfinterfaces.IFlag{}
	//cmd :=
	shell.RegisterNewCommand("env", "Environment command", Env)
	shell.RegisterCommandNewStringFlag("env", "key", "", "Sets a string value")
	shell.RegisterCommandNewStringFlag("env", "value", "", "Sets a string value")
	shell.RegisterCommandNewBoolFlag("env", "list", false, "List Environment Variables")

	//Flags.NewStringFlag("key", "", "Sets a string value")
	//Flags.NewStringFlag("value", "", "Sets a string value")
	//Flags.NewBoolFlag("list", false, "List Environment Variables")

	//shell.AddNewCommandWithFlags("env", "Environment command", Env, Flags)
}

func Env(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()

	fg := command.GetCommandFlags()
	ci := command.GetCommandInput()
	args := ci.GetArgs()

	pflags := fg.Parsedflags()
	env := shell.GetEnvironment()

	kflag := pflags["key"]
	vflag := pflags["value"]

	log.LogDebugf("env()", "key GetName = %s", kflag.GetName())
	log.LogDebugf("env()", "value GetName = %s", vflag.GetName())

	kstr := kflag.GetStringValue()
	vstr := vflag.GetStringValue()

	log.LogDebugf("env()", "kstr = %s", *kstr)
	log.LogDebugf("env()", "vstr = %s", *vstr)

	if *kstr != kflag.GetDefaultStringValue() {

		if *vstr != vflag.GetDefaultStringValue() && args[2] == "set" {
			shell.Println("--")
			shell.Printlnf("Setting %s to %s", *kstr, *vstr)
			shell.Println("--")
			env.Set(env.MakeSingleVariable(*kstr, *vstr))
		} else {
			if args[2] == "clear" {
				shell.Println("--")
				shell.Printlnf("Clearing %s", *kstr)
				shell.Println("--")
				env.Clear(*kstr)
			}

			if args[2] == "delete" {
				shell.Println("--")
				shell.Printlnf("Deleting %s", *kstr)
				shell.Println("--")
				env.Delete(*kstr)
			}
		}

	} else {
		log.LogDebugf("env()", "skipping: kstr = %s matched default %s", *kstr, kflag.GetDefaultStringValue())
		log.LogDebugf("env()", "skipping: vstr = %s macthed default %s", *vstr, vflag.GetDefaultStringValue())
	}

	listflag := pflags["list"]

	if listflag != nil && listflag.GetBoolValue() != nil && *listflag.GetBoolValue() {
		log.LogDebugf("env()", "listflag is not nil")

		namevalues := env.GetNameValues()
		for k, _ := range namevalues {
			envp := namevalues[k]
			shell.Printlnf("%s = %s", k, strings.Join(envp.GetValues(), ","))
		}
	}

	return command.NewSuccessCommandResult("")
}
