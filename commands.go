package shellframework

import (
	"strings"

	"github.com/eshu0/shellframework/interfaces"
)

type Command struct {
	name         string
	description  string
	operator     func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult
	shell        sfinterfaces.IShell
	commandinput sfinterfaces.ICommandInput
	flags        sfinterfaces.IFlags
}

func (command *Command) GetDescription() string {
	return command.description
}

func (command *Command) GetName() string {
	return command.name
}

func (command *Command) SetShell(shell sfinterfaces.IShell) {
	command.shell = shell
}

func (command *Command) GetShell() sfinterfaces.IShell {
	return command.shell
}

//Get and Set the input for the command
func (command *Command) SetCommandInput(input sfinterfaces.ICommandInput) {
	command.commandinput = input
}

func (command *Command) GetCommandInput() sfinterfaces.ICommandInput {
	return command.commandinput
}

// This matches the command
func (command *Command) Match(incmd sfinterfaces.ICommandInput) bool {

	shell := command.GetShell()
	log := *shell.GetLog()

	if strings.HasPrefix(incmd.GetLowerCommandName(), strings.ToLower(command.GetName())) {
		log.LogDebug("Match()", "Command '%s' matched '%s'", command.GetName(), incmd.GetLowerCommandName())
		return true
	}

	// this will space the trace with false positives
	//log.LogDebug("Match()", "Command '%s' did not match '%s'", command.GetName(), incmd.GetLowerCommandName())
	return false
}

func (command *Command) Process() sfinterfaces.ICommandResult {

	// get the shell for logging
	shell := command.GetShell()
	log := *shell.GetLog()

	ci := command.GetCommandInput()

	args := ci.GetArgs()
	log.LogDebug("Process()", "Number of args: %d", len(args))

	flgs := command.GetFlags()

	log.LogDebug("Process()", "Parsing the flags")
	flgs.Parse()

	parsedflags := flgs.Parsedflags()

	log.LogDebug("Process()", "Number of args: %d", len(args))

	// have to derefence due to the interface
	flgset := flgs.GetFlagSet()

	if flgset != nil {
		log.LogDebug("Process()", "flagset not nil - Parsing flagset")
		flgset.Parse(args)

		for _, arg := range flgset.Args() {
			log.LogDebug("Process()", "Argument: %s", arg)
		}
	} else {
		log.LogDebug("Process()", "flagset was nil ")
	}

	for _, sflag := range parsedflags {
		log.LogDebug("Process()", "Parsed Flag - GetName %s", sflag.GetName())
		log.LogDebug("Process()", "Parsed Flag - GetStringValue %s", sflag.GetStringValue())
		log.LogDebug("Process()", "Parsed Flag - GetBoolValue %s", sflag.GetBoolValue())
		log.LogDebug("Process()", "Parsed Flag - GetIntValue %s", sflag.GetIntValue())
		log.LogDebug("Process()", "Parsed Flag - GetFlagType %s", sflag.GetFlagType())

	}

	log.LogDebug("Process()", "running command %s", command.GetName())
	result := command.operator(command)
	log.LogDebug("Process()", "finished command %s", command.GetName())
	log.LogDebug("Process()", "result - Successful %t", result.Sucessful())
	log.LogDebug("Process()", "result - ExitShell %t", result.ExitShell())
	log.LogDebug("Process()", "result - Error %s", result.Err())
	log.LogDebug("Process()", "result - Result %s", result.Result())
	return result
}

// This will regsiter the command with the shell
// If adding commands then implement this method
func (command *Command) Register(shell sfinterfaces.IShell) {

}

func (command *Command) GetFlags() sfinterfaces.IFlags {
	return command.flags
}

func (command *Command) SetFlags(flgs sfinterfaces.IFlags) {
	command.flags = flgs
}

func (command *Command) NewSuccessCommandResult(msg string) sfinterfaces.ICommandResult {
	//var sr CommandResult
	sr := &CommandResult{}
	sr.err = nil
	sr.sucess = true
	sr.result = msg
	sr.exitshell = false
	return sr
}

func (command *Command) NewExitSuccessCommandResult() sfinterfaces.ICommandResult {
	//var sr CommandResult
	sr := &CommandResult{}
	sr.err = nil
	sr.sucess = true
	sr.result = ""
	sr.exitshell = true
	return sr
}

type CommandInput struct {
	name     string
	rawinput string
	args     []string
}

func (cinput *CommandInput) GetCommandName() string {
	return cinput.name
}

func (cinput *CommandInput) GetLowerCommandName() string {
	return strings.ToLower(cinput.name)
}

func (cinput *CommandInput) GetArgs() []string {
	return cinput.args
}

func (cinput *CommandInput) GetRawInput() string {
	return cinput.rawinput
}

func (cinput *CommandInput) GetInputWithOutCommand() string {
	if cinput.rawinput != "" {
		runes := []rune(cinput.rawinput)
		commandlength := len(cinput.name)
		return string(runes[commandlength:])
	}
	return ""
}

func (cinput *CommandInput) SetCommandName(name string) {
	cinput.name = name
}

func (cinput *CommandInput) SetRawInput(input string) {
	cinput.rawinput = input
}

func (cinput *CommandInput) SetArgs(args []string) {
	cinput.args = args
}

type CommandResult struct {
	exitshell bool
	sucess    bool
	err       error
	result    string
}

func (cresult *CommandResult) ExitShell() bool {
	return cresult.exitshell
}

func (cresult *CommandResult) Sucessful() bool {
	return cresult.sucess
}

func (cresult *CommandResult) Err() error {
	return cresult.err
}

func (cresult *CommandResult) Result() string {
	return cresult.result
}
