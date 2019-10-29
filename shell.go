package shellframework

import (
	"os"

	"github.com/eshu0/shellframework/interfaces"
)

type Shell struct {
	commands    []sfinterfaces.ICommand
	environment *sfinterfaces.IEnvironment
	version     string
	session     *sfinterfaces.ISession

	log *sfinterfaces.IShellLogger
	in  *os.File
	out *os.File
	err *os.File
}

func NewShell(session sfinterfaces.ISession, In *os.File, Out *os.File, Err *os.File, logfile sfinterfaces.IShellLogger) *Shell {

	shell := &Shell{}

	env := NewEnvironment(shell)
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

func (shell *Shell) GetCommands() []sfinterfaces.ICommand {
	return shell.commands
}

func (shell *Shell) GetEnvironment() sfinterfaces.IEnvironment {
	return *shell.environment
}

func (shell *Shell) GetSession() sfinterfaces.ISession {
	return *shell.session
}

func (shell *Shell) GetVersion() string {
	return shell.version
}

func (shell *Shell) GetIn() *os.File {
	return shell.in
}

func (shell *Shell) GetOut() *os.File {
	return shell.out
}

func (shell *Shell) GetErr() *os.File {
	return shell.err
}

//
// SHELL Logging
//
// these function provide logging to the choosen logfile
//

func (shell *Shell) SetLog(logger sfinterfaces.IShellLogger) {
	shell.log = &logger
}

func (shell *Shell) GetLog() *sfinterfaces.IShellLogger {
	return shell.log
}
