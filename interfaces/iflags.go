package sfinterfaces

import "flag"

type IFlags interface {

	// what was read from the shell
	//SetFlagValue(toread *flag.FlagSet, flg IFlag)
	GetFlagSet() *flag.FlagSet

	//Get and Set IExecution
	SetCommand(input ICommand)
	GetCommand() ICommand

	// Argument Parsing from this point
	Parsedflags() map[string]IFlag

	GetFlags() []IFlag
	SetFlags(flgs []IFlag)

	Parse()

	NewIntFlag(name string, defaultvalue int, usage string)
	NewBoolFlag(name string, defaultvalue bool, usage string)
	NewStringFlag(name string, defaultvalue string, usage string)
}

// I'd like to replace this with the flags
// however flags assumes the os.Args() which we are not doing
// we are well beyond the processing of the Args
// if i find something later this will die a death
type IFlag interface {
	// name of the flag
	GetName() string

	// what we are expecting
	// three supported at time of wrtiing
	// int, string and bool
	GetFlagType() int
	// usage for the flag
	// this is the same as flag interface
	GetUsage() string

	// defaults
	GetDefaultBoolValue() bool
	GetDefaultStringValue() string
	GetDefaultIntValue() int

	// what was read from the shell
	GetStringValue() *string
	GetBoolValue() *bool
	GetIntValue() *int

	// used for reading from the flag set
	SetFlagValue(toread *flag.FlagSet)
}
