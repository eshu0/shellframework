package sfinterfaces

import (
	"os"
	"github.com/eshu0/simplelogger/interfaces"
)

// this is the Shell version
const Version = "0.1"

// main interface for the Shell
// i am human so most likely i have forgetten something
type IShell interface {

	// getters and setters
	GetCommands() []ICommand
	SetCommands(cmds []ICommand)

	GetSession() ISession
	SetSession(session ISession)

	GetEnvironment() IEnvironment
	SetEnvironment(env IEnvironment)

	GetVersion() string

	// logging
	GetLog() *slinterfaces.ISimpleLogger
	SetLog(log slinterfaces.ISimpleLogger)

	// in and out
	GetIn() *os.File
	GetOut() *os.File
	GetErr() *os.File

	//methods run the shell
	Run()
	NonInteractiveSession(env IEnvironment, log  slinterfaces.ISimpleLogger)
	InteractiveSession(env IEnvironment, log  slinterfaces.ISimpleLogger)

	ParseInput(input string) []ICommandInput

	// print functions
	Println(msg string)
	Printlnf(msg string, a ...interface{})

	// extra print functions
	PrintDetails()
	// called during the run command
	PrintInputMessage()

	// add commands to command or commands to the list
	AddCommand(cmd ICommand)
	AddCommands(commands []ICommand)

	// create a new command
	NewCommand(name string, description string, operator func(command ICommand) ICommandResult, flags []IFlag) ICommand

	// register a new command
	RegisterNewCommand(name string, description string, operator func(command ICommand) ICommandResult)
	RegisterNewCommandWithFlags(name string, description string, operator func(command ICommand) ICommandResult, flags []IFlag)

	// register a single flag to a command
	RegisterCommandNewIntFlag(cmd string, name string, defaultvalue int, usage string)
	RegisterCommandNewBoolFlag(cmd string, name string, defaultvalue bool, usage string)
	RegisterCommandNewStringFlag(cmd string, name string, defaultvalue string, usage string)

	RegisterCommandFlag(cmd string, flag IFlag)
}
