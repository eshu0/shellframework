package simpleshell

import (
	"os"

	"github.com/eshu0/shellframework"
)

type SimpleShell struct {
	commands    []shellframework.ICommand
	environment *shellframework.IEnvironment
	version     string
	session     *shellframework.ISession

	log *shellframework.IShellLogger
	in  *os.File
	out *os.File
	err *os.File
}

func NewSimpleShell(session shellframework.ISession, In *os.File, Out *os.File, Err *os.File, logfile shellframework.IShellLogger) *SimpleShell {

	shell := &SimpleShell{}

	env := NewSimpleEnvironment(shell)
	shell.environment = &env

	shell.commands = []shellframework.ICommand{}
	shell.version = shellframework.Version
	shell.session = &session

	shell.in = In
	shell.out = Out
	shell.err = Err
	shell.log = &logfile

	return shell
}

// Get Methods

func (shell *SimpleShell) GetCommands() []shellframework.ICommand {
	return shell.commands
}

func (shell *SimpleShell) GetEnvironment() shellframework.IEnvironment {
	return *shell.environment
}

func (shell *SimpleShell) GetSession() shellframework.ISession {
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

func (shell *SimpleShell) SetLog(logger shellframework.IShellLogger) {
	shell.log = &logger
}

func (shell *SimpleShell) GetLog() *shellframework.IShellLogger {
	return shell.log
}
