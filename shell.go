package simpleshell

import (
	"os"

	"github.com/eshu0/shellframework/interfaces"
)

type SimpleShell struct {
	commands    []sfinterfaces.ICommand
	environment *sfinterfaces.IEnvironment
	version     string
	session     *sfinterfaces.ISession

	log *sfinterfaces.IShellLogger
	in  *os.File
	out *os.File
	err *os.File
}

func NewSimpleShell(session sfinterfaces.ISession, In *os.File, Out *os.File, Err *os.File, logfile sfinterfaces.IShellLogger) *SimpleShell {

	shell := &SimpleShell{}

	env := NewSimpleEnvironment(shell)
	shell.environment = &env

	shell.commands = []sfinterfaces.ICommand{}
	shell.version = sfinterfaces.Version
	shell.session = &session

	shell.in = In
	shell.out = Out
	shell.err = Err
	shell.log = &logfile

	return shell
}

// Get Methods

func (shell *SimpleShell) GetCommands() []sfinterfaces.ICommand {
	return shell.commands
}

func (shell *SimpleShell) GetEnvironment() sfinterfaces.IEnvironment {
	return *shell.environment
}

func (shell *SimpleShell) GetSession() sfinterfaces.ISession {
	return *shell.session
}

func (shell *SimpleShell) GetVersion() string {
	return shell.version
}

func (shell *SimpleShell) GetIn() *os.File {
	return shell.in
}

func (shell *SimpleShell) GetOut() *os.File {
	return shell.out
}

func (shell *SimpleShell) GetErr() *os.File {
	return shell.err
}

//
// SHELL Logging
//
// these function provide logging to the choosen logfile
//

func (shell *SimpleShell) SetLog(logger sfinterfaces.IShellLogger) {
	shell.log = &logger
}

func (shell *SimpleShell) GetLog() *sfinterfaces.IShellLogger {
	return shell.log
}
