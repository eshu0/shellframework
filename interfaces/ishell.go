package sfinterfaces

import (
	"os"
)

// this is the Shell version
const Version = "0.1"

// main interface for the Shell
// i am human so most likely i have forgetten something
type IShell interface {

	// getters
	GetCommands() []ICommand
	GetEnvironment() IEnvironment
	GetSession() ISession
	GetVersion() string

	// logging
	GetLog() *IShellLogger
	SetLog(log IShellLogger)

	// in and out
	GetIn() *os.File
	GetOut() *os.File
	GetErr() *os.File

	//methods run the shell
	Run()
	ParseInput(input string) []ICommandInput

	// print functions
	Println(msg string)
	Printlnf(msg string, a ...interface{})

	// extra print functions
	PrintDetails()
	// called during the run command
	PrintInputMessage()

	AddCommand(cmd ICommand)
	AddCommands(commands []ICommand)

	NewCommand(name string, description string, operator func(command ICommand) ICommandResult, flags []IFlag) ICommand
	AddNewCommand(name string, description string, operator func(command ICommand) ICommandResult)
	AddNewCommandWithFlags(name string, description string, operator func(command ICommand) ICommandResult, flags []IFlag)
}
