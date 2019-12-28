package shellframework

import (
	"flag"
	"fmt"

	"github.com/eshu0/shellframework/interfaces"
)

type CommandFlags struct {
	sfinterfaces.IFlags
	command     sfinterfaces.ICommand
	parsedflags map[string]sfinterfaces.IFlag
	flags       []sfinterfaces.IFlag
}

type CommandFlag struct {
	sfinterfaces.IFlag
	name  string
	usage string

	defaultbvalue bool
	defaultsvalue string
	defaultivalue int

	flagtype int

	foundbvalue *bool
	foundsvalue *string
	foundivalue *int
}

func (flg *CommandFlag) String() string {
	switch flg.GetFlagType() {
	case 1:
		if flg.GetStringValue() == nil {
			return ""
		} else {
			return *flg.GetStringValue()
		}
	case 2:
		if flg.GetBoolValue() == nil {
			return ""
		} else {
			return fmt.Sprintf("%t", *flg.GetBoolValue())
		}
	case 3:
		if flg.GetIntValue() == nil {
			return ""
		} else {
			return fmt.Sprintf("%d", *flg.GetIntValue())
		}
	}

	return "Unknown Type"
}

func (flg *CommandFlag) GetName() string {
	return flg.name
}

func (flg *CommandFlag) GetFlagType() int {
	return flg.flagtype
}

func (flg *CommandFlag) GetUsage() string {
	return flg.usage
}

func (flg *CommandFlag) GetDefaultBoolValue() bool {
	return flg.defaultbvalue
}

func (flg *CommandFlag) GetDefaultStringValue() string {
	return flg.defaultsvalue
}
func (flg *CommandFlag) GetDefaultIntValue() int {
	return flg.defaultivalue
}

func (flg *CommandFlag) GetStringValue() *string {
	return flg.foundsvalue
}
func (flg *CommandFlag) GetBoolValue() *bool {
	return flg.foundbvalue
}
func (flg *CommandFlag) GetIntValue() *int {
	return flg.foundivalue
}

func (flg *CommandFlag) SetFlagValue(toread *flag.FlagSet) {

	//flgs.parent.flagset = toread

	if toread != nil {

		switch flg.GetFlagType() {
		case 1:
			//toread.StringVar(flg.GetStringValue(), flg.GetName(), flg.GetDefaultStringValue(), flg.GetUsage())
			flg.foundsvalue = toread.String(flg.GetName(), flg.GetDefaultStringValue(), flg.GetUsage())
		case 2:
			//toread.BoolVar(flg.GetBoolValue(), flg.GetName(), flg.GetDefaultBoolValue(), flg.GetUsage())
			flg.foundbvalue = toread.Bool(flg.GetName(), flg.GetDefaultBoolValue(), flg.GetUsage())
		case 3:
			//toread.IntVar(flg.GetIntValue(), flg.GetName(), flg.GetDefaultIntValue(), flg.GetUsage())
			flg.foundivalue = toread.Int(flg.GetName(), flg.GetDefaultIntValue(), flg.GetUsage())

		}
	} else {

	}

}

// flags after here

func (sflgs *CommandFlags) Parsedflags() map[string]sfinterfaces.IFlag {
	return sflgs.parsedflags
}

func (sflgs *CommandFlags) GetFlags() []sfinterfaces.IFlag {
	return sflgs.flags
}

func (sflgs *CommandFlags) SetFlags(flgs []sfinterfaces.IFlag) {
	sflgs.flags = flgs
}

func (flgs *CommandFlags) SetCommand(cmd sfinterfaces.ICommand) {
	flgs.command = cmd
}

func (flgs *CommandFlags) GetCommand() sfinterfaces.ICommand {
	return flgs.command
}

func (flgs *CommandFlags) PrintUsage() {

	command := flgs.GetCommand()

	//get she;;
	shell := command.GetShell()
	log := *shell.GetLog()
	// flags
	flags := flgs.GetFlags()
	log.LogDebugf("PrintUsage()", " Number of flags %d for %s ", len(flags), command.GetName())

	for _, flg := range flags {
		shell.Println("----")
		shell.Printlnf("-%s", flg.GetName())
		shell.Printlnf("\t Type %d", flg.GetFlagType())
		shell.Println("\t Usage:")
		shell.Printlnf("\t\t%s", flg.GetUsage())
		shell.Println("----")

		log.LogDebugf("PrintUsage()", "%s %s", flg.GetName(), flg)

	}

}

func (flgs *CommandFlags) Parse() {

	command := flgs.GetCommand()
	shell := command.GetShell()
	log := *shell.GetLog()

	// flags
	flags := flgs.GetFlags()

	// parsed flags
	flgs.parsedflags = make(map[string]sfinterfaces.IFlag)

	flgset := flag.NewFlagSet(command.GetName(), flag.ContinueOnError)

	log.LogDebugf("Parse()", " Number of flags %d for %s ", len(flags), command.GetName())

	for _, flg := range flags {

		//fr := &FlagResult{}
		//fr.SetFlag(flg)
		//_, alreadythere := sc.formal[flg.name]
		log.LogDebugf("Parse()", "Look up flag : %s ", flg.GetName())
		alreadythere := flgset.Lookup(flg.GetName())
		if alreadythere == nil {
			log.LogDebugf("Parse()", "%s was nil which means it is missing so going to add", flg.GetName())
			flg.SetFlagValue(flgset)
			flgs.parsedflags[flg.GetName()] = flg
		} else {
			log.LogDebugf("Parse()", "%s was not nil so it will not be added", flg.GetName())
		}
	}

	ci := command.GetCommandInput()

	if ci != nil {
		args := ci.GetArgs()
		if args != nil {
			flgset.Parse(args)
		} else {
			log.LogDebug("Parse()", "args was nil ")
		}
	} else {
		log.LogDebug("Parse()", "ci was nil ")
	}

	/*
		flgset.VisitAll(func(f *flag.Flag) {
			log.LogDebugf("Parse()", "VisitAll - %s %s %s", f.Value, f.Name, f.Usage)
			flag.Var(f.Value, f.Name, f.Usage)
		})
	*/

	log.LogDebug("Parse()", "Set the parsed flags")

	for _, sflag := range flgs.parsedflags {
		log.LogDebugf("Parse()", "flag.GetName: %s", sflag.GetName())
		log.LogDebugf("Parse()", "flag.GetFlagType: %d", sflag.GetFlagType())
		log.LogDebugf("Parse()", "flag %s", sflag)

	}

}
