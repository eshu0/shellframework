package simpleshell

import (
	"flag"

	"github.com/eshu0/shellframework"
)

type SimpleFlags struct {
	flagset     *flag.FlagSet
	command     shellframework.ICommand
	parsedflags map[string]shellframework.IFlag
	flags       []shellframework.IFlag
}

type SimpleFlag struct {
	name  string
	usage string

	defaultbvalue bool
	defaultsvalue string
	defaultivalue int

	flagtype int

	foundbvalue *bool
	foundsvalue *string
	foundivalue *int

	//parent SimpleFlags
}

func NewBoolFlag(name string, defaultvalue bool, usage string) shellframework.IFlag {
	sf := &SimpleFlag{}
	sf.name = name
	sf.defaultbvalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 2
	return sf
}

func NewIntFlag(name string, defaultvalue int, usage string) shellframework.IFlag {
	sf := &SimpleFlag{}
	sf.name = name
	sf.defaultivalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 3
	return sf
}

func NewStringFlag(name string, defaultvalue string, usage string) shellframework.IFlag {
	sf := &SimpleFlag{}
	sf.name = name
	sf.defaultsvalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 1
	return sf
}

func (flg *SimpleFlag) GetName() string {
	return flg.name
}

func (flg *SimpleFlag) GetFlagType() int {
	return flg.flagtype
}

func (flg *SimpleFlag) GetUsage() string {
	return flg.usage
}

func (flg *SimpleFlag) GetDefaultBoolValue() bool {
	return flg.defaultbvalue
}

func (flg *SimpleFlag) GetDefaultStringValue() string {
	return flg.defaultsvalue
}
func (flg *SimpleFlag) GetDefaultIntValue() int {
	return flg.defaultivalue
}

func (flg *SimpleFlag) GetStringValue() *string {
	return flg.foundsvalue
}
func (flg *SimpleFlag) GetBoolValue() *bool {
	return flg.foundbvalue
}
func (flg *SimpleFlag) GetIntValue() *int {
	return flg.foundivalue
}

func (flg *SimpleFlag) SetFlagValue(toread *flag.FlagSet) {

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

func (sflgs *SimpleFlags) Parsedflags() map[string]shellframework.IFlag {
	return sflgs.parsedflags
}

func (sflgs *SimpleFlags) GetFlags() []shellframework.IFlag {
	return sflgs.flags
}

func (sflgs *SimpleFlags) SetFlags(flgs []shellframework.IFlag) {
	sflgs.flags = flgs
}

func (flgs *SimpleFlags) GetFlagSet() *flag.FlagSet {
	return flgs.flagset
}

func (flgs *SimpleFlags) SetCommand(cmd shellframework.ICommand) {
	flgs.command = cmd
}

func (flgs *SimpleFlags) GetCommand() shellframework.ICommand {
	return flgs.command
}

func (flgs *SimpleFlags) Parse() {

	command := flgs.GetCommand()

	//get she;;
	shell := command.GetShell()
	log := *shell.GetLog()

	flgset := flag.NewFlagSet(command.GetName(), flag.ContinueOnError)

	// flags
	flags := flgs.GetFlags()

	// parsed flags
	flgs.parsedflags = make(map[string]shellframework.IFlag)

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
