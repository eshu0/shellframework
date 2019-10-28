package shellframework

import (
	"strings"

	"github.com/eshu0/shellframework/interfaces"
)

type SimpleCommand struct {
	name        string
	description string
	operator    func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult
	shell       sfinterfaces.IShell
	//args         []string
	commandinput sfinterfaces.ICommandInput
	flags        sfinterfaces.IFlags
}

func (command *SimpleCommand) GetDescription() string {
	return command.description
}

func (command *SimpleCommand) GetName() string {
	return command.name
}

func (command *SimpleCommand) SetShell(shell sfinterfaces.IShell) {
	command.shell = shell
}

func (command *SimpleCommand) GetShell() sfinterfaces.IShell {
	return command.shell
}

//Get and Set the input for the command
func (command *SimpleCommand) SetCommandInput(input sfinterfaces.ICommandInput) {
	command.commandinput = input
}

func (command *SimpleCommand) GetCommandInput() sfinterfaces.ICommandInput {
	return command.commandinput
}

// This matches the command
func (command *SimpleCommand) Match(incmd sfinterfaces.ICommandInput) bool {

	shell := command.GetShell()
	log := *shell.GetLog()

	if strings.HasPrefix(incmd.GetLowerCommandName(), strings.ToLower(command.GetName())) {
		log.LogPrintlnf("Match(): Command '%s' matched '%s'", command.GetName(), incmd.GetLowerCommandName())
		return true
	}

	// this will space the trace with false positives
	//
	//command.shell.LogPrintlnf("Match(): Command '%s' did not match '%s'", command.GetName(), incmd.GetLowerCommandName())
	return false
}

func (command *SimpleCommand) Process() sfinterfaces.ICommandResult {
	// get the shell for logging

	shell := command.GetShell()
	log := *shell.GetLog()
	/*
		ci := command.GetCommandInput()

		args := ci.GetArgs()
		shell.LogPrintlnf("Process(): Number of args: %d", len(args))
	*/

	flgs := command.GetFlags()

	log.LogPrintln("Process(): Parsing the flags")
	flgs.Parse()
	/*


		//flgset := carg.GetParsedFlagSet()
		parsedflags := flgs.Parsedflags()

		shell.LogPrintlnf("Process(): Number of args: %d", len(args))

		if args != nil && len(args) > 0 && parsedflags != nil {

			for _, sflag := range parsedflags {
				// have to derefence due to the interface
				//sflag := *p
				flgset := flgs.GetFlagSet()

				if flgset != nil {
					shell.LogPrintln("Process(): flagset not nil - Parsing flagset")
					flgset.Parse(args)

					for _, arg := range flgset.Args() {
						shell.LogPrintlnf("Process(): Argument: %s", arg)
					}
				} else {
					shell.LogPrintln("Process(): flagset was nil ")
				}
			}

		} else {
			shell.LogPrintln("Process(): Not parsing flagset")
		}
	*/
	return command.operator(command)
}

/*
func (command *SimpleCommand) Args() []string {
	return command.args
}
*/

func (command *SimpleCommand) GetFlags() sfinterfaces.IFlags {
	return command.flags
}

func (command *SimpleCommand) SetFlags(flgs sfinterfaces.IFlags) {
	command.flags = flgs
}

type SimpleCommandInput struct {
	name     string
	rawinput string
	args     []string
}

func (cinput *SimpleCommandInput) GetCommandName() string {
	return cinput.name
}

func (cinput *SimpleCommandInput) GetLowerCommandName() string {
	return strings.ToLower(cinput.name)
}

func (cinput *SimpleCommandInput) GetArgs() []string {
	return cinput.args
}

func (cinput *SimpleCommandInput) GetRawInput() string {
	return cinput.rawinput
}

func (cinput *SimpleCommandInput) GetInputWithOutCommand() string {
	if cinput.rawinput != "" {
		runes := []rune(cinput.rawinput)
		commandlength := len(cinput.name)
		return string(runes[commandlength:])
	}
	return ""
}

func (cinput *SimpleCommandInput) SetCommandName(name string) {
	cinput.name = name
}

func (cinput *SimpleCommandInput) SetRawInput(input string) {
	cinput.rawinput = input
}

func (cinput *SimpleCommandInput) SetArgs(args []string) {
	cinput.args = args
}

type SimpleCommandResult struct {
	exitshell bool
	sucess    bool
	err       error
	result    string
}

func (cresult *SimpleCommandResult) ExitShell() bool {
	return cresult.exitshell
}

func (cresult *SimpleCommandResult) Sucessful() bool {
	return cresult.sucess
}

func (cresult *SimpleCommandResult) Err() error {
	return cresult.err
}

func (cresult *SimpleCommandResult) Result() string {
	return cresult.result
}

func NewSuccessCommandResult(msg string) sfinterfaces.ICommandResult {
	//var sr SimpleCommandResult
	sr := &SimpleCommandResult{}
	sr.err = nil
	sr.sucess = true
	sr.result = msg
	sr.exitshell = false
	return sr
}

func NewExitSuccessCommandResult() sfinterfaces.ICommandResult {
	//var sr SimpleCommandResult
	sr := &SimpleCommandResult{}
	sr.err = nil
	sr.sucess = true
	sr.result = ""
	sr.exitshell = true
	return sr
}
