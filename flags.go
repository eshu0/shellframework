package shellframework

import (
	"flag"

	"github.com/eshu0/shellframework/interfaces"
)

type CommandFlags struct {
	flagset     *flag.FlagSet
	command     sfinterfaces.ICommand
	parsedflags map[string]sfinterfaces.IFlag
	flags       []sfinterfaces.IFlag
}

type CommandFlag struct {
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

func (flgs *CommandFlags) GetFlagSet() *flag.FlagSet {
	return flgs.flagset
}

func (flgs *CommandFlags) SetCommand(cmd sfinterfaces.ICommand) {
	flgs.command = cmd
}

func (flgs *CommandFlags) GetCommand() sfinterfaces.ICommand {
	return flgs.command
}

func (flgs *CommandFlags) Parse() {

	command := flgs.GetCommand()

	//get she;;
	shell := command.GetShell()
	log := *shell.GetLog()

	flgset := flag.NewFlagSet(command.GetName(), flag.ContinueOnError)

	// flags
	flags := flgs.GetFlags()

	// parsed flags
	flgs.parsedflags = make(map[string]sfinterfaces.IFlag)

	log.LogPrintlnf("Parse(): Number of flags %d for %s ", len(flags), command.GetName())

	for _, flg := range flags {

		//fr := &FlagResult{}
		//fr.SetFlag(flg)
		//_, alreadythere := sc.formal[flg.name]
		log.LogPrintlnf("Parse(): Look up flag : %s ", flg.GetName())
		alreadythere := flgset.Lookup(flg.GetName())
		if alreadythere == nil {
			log.LogPrintlnf("Parse(): %s was nil which means it is missing so going to add", flg.GetName())
			flg.SetFlagValue(flgset)
			flgs.parsedflags[flg.GetName()] = flg
		} else {
			log.LogPrintlnf("Parse(): %s was not nil so it will not be added", flg.GetName())
		}

	}

	//command.parsedflags = parsedflags
	log.LogPrintln("Parse(): Set the parsed flags")

	//if command.Shell.env_trace >= 0 {

	for _, sflag := range flgs.parsedflags {
		// have to derefence due to the interface
		//sflag := *p
		// what type are we dealing with
		log.LogPrintlnf("Parse(): flag.GetName: %s", sflag.GetName())
		log.LogPrintlnf("Parse(): flag.GetFlagType: %d", sflag.GetFlagType())

		switch sflag.GetFlagType() {
		case 1:
			log.LogPrintlnf("Parse(): flag.String: %s", sflag.GetStringValue())
		case 2:
			log.LogPrintlnf("Parse(): flag.Boolean: %t", sflag.GetBoolValue())
		case 3:
			log.LogPrintlnf("Parse(): flag.Integer: %d", sflag.GetIntValue())
		}

	}

}
