package defaultcmds

import (
	"strings"

	"github.com/eshu0/shellframework"
	"github.com/eshu0/shellframework/interfaces"
)

func env(command sfinterfaces.ICommand) sfinterfaces.ICommandResult {

	shell := command.GetShell()
	log := *shell.GetLog()

	fg := command.GetFlags()

	pflags := fg.Parsedflags()
	env := shell.GetEnvironment()

	log.LogPrintln("env() command called")
	/*
		for _, sflag := range pflags {
			// have to derefence due to the interface
			//sflag := *p
			shell.LogPrintlnf("env(): Argument: %s", sflag.GetName())
			flgset := sflag.GetFlagSet()

			if flgset != nil {
				shell.LogPrintln("env(): flagset not nil")
				for _, arg := range flgset.Args() {
					shell.LogPrintlnf("env(): Argument: %s", arg)
				}
			} else {
				shell.LogPrintln("env(): flagset was nil ")
			}
		}
	*/
	// map to [string]*IEnvironmentVariable
	kflag := pflags["key"]
	vflag := pflags["value"]

	//if kflag != nil && vflag != nil {

	//shell.LogPrintln("env() kflag and vflag is not nil")

	//keyflag := *kflag
	//valueflag := *vflag

	log.LogPrintlnf("env() key GetName = %s", kflag.GetName())
	log.LogPrintlnf("env() value GetName = %s", vflag.GetName())

	//if keyflag != nil && valueflag != nil {
	//shell.LogPrintln("env() keyflag and valueflag is not nil")

	kstr := kflag.GetStringValue()
	vstr := vflag.GetStringValue()

	log.LogPrintlnf("env() kstr = %s", *kstr)
	log.LogPrintlnf("env() vstr = %s", *vstr)

	//if len(arguements.Args) >= 0 {
	//if keyflag != nil && kstr != nil && *kstr != "" && valueflag != nil && vstr != nil && *vstr != "" {
	if *kstr != kflag.GetDefaultStringValue() && *vstr != vflag.GetDefaultStringValue() {

		//if arguements.Args[0] == "set" {
		shell.Println("--")
		shell.Printlnf("Setting %s to %s", *kstr, *vstr)
		shell.Println("--")
		env.SetVariable(env.MakeSingleVariable(*kstr, *vstr))
		//}
	} else {
		log.LogPrintlnf("env() skipping: kstr = %s matched default %s", *kstr, kflag.GetDefaultStringValue())
		log.LogPrintlnf("env() skipping: vstr = %s macthed default %s", *vstr, vflag.GetDefaultStringValue())
	}

	//}
	//}
	//}

	listflag := pflags["list"]
	//lbool := listflag.GetBoolValue()
	if listflag != nil { //&& lbool == true {
		log.LogPrintln("env() listflag is not nil")

		//	if listflag != nil && lbool != nil && *lbool == true {
		namevalues := env.GetNameValues()
		for k, _ := range namevalues {
			envp := namevalues[k]
			shell.Printlnf("%s = %s", k, strings.Join(envp.GetValues(), ","))
		}
	}

	return shellframework.NewSuccessCommandResult("")
}
