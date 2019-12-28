package shellframework

import (
	"os"

	"github.com/eshu0/shellframework/defaultcmds"
	"github.com/eshu0/shellframework/interfaces"
)

type Shell struct {
	sfinterfaces.IShell
	commands    []sfinterfaces.ICommand
	environment *sfinterfaces.IEnvironment
	version     string
	session     *sfinterfaces.ISession

	log *sfinterfaces.IShellLogger
	in  *os.File
	out *os.File
	err *os.File
}

func NewShellFromFile(session sfinterfaces.ISession,scriptinfilepath string, Out *os.File, Err *os.File, logi sfinterfaces.IShellLogger) *Shell {

	var input *os.File

	// script file been provided as input to the cli?
	if(scriptinfilepath != "") {

		// open the file if we cannot then crash out
		inf, err := os.Open(scriptinfilepath)
		if err != nil {
			// should write an error here
			return nil
		}

		// replace the Stdin with the file we have read
		input = inf

		// we don't want an interactive session as we are running a script
		session.SetInteractive(false)
	} else {
		// should write an error here
		return nil
	}

	shell := NewShell(session,input, Out, Err, logi)
	return shell
}

func NewShell(session sfinterfaces.ISession, In *os.File, Out *os.File, Err *os.File, logi sfinterfaces.IShellLogger) *Shell {

	shell := &Shell{}

	env := NewEnvironment(shell)
	shell.SetEnvironment(env)
	shell.SetSession(session)

	shell.version = sfinterfaces.Version

	shell.in = In
	shell.out = Out
	shell.err = Err
	shell.log = &logi

	shell.commands = []sfinterfaces.ICommand{}

	// Add the default commands
	mcmd := dcmds.ManCommand{}
	mcmd.Register(shell)

	hcmd := dcmds.HelpCommand{}
	hcmd.Register(shell)

	excmd := dcmds.ExitCommand{}
	excmd.Register(shell)

	ccmd := dcmds.CmdCommand{}
	ccmd.Register(shell)

	envcmd := dcmds.EnvCommand{}
	envcmd.Register(shell)

	echocmd := dcmds.EchoCommand{}
	echocmd.Register(shell)

	return shell
}

// Get Methods

func (shell *Shell) SetCommands(cmds []sfinterfaces.ICommand) {
	shell.commands = cmds
}

func (shell *Shell) GetCommands() []sfinterfaces.ICommand {
	return shell.commands
}

func (shell *Shell) SetEnvironment(env sfinterfaces.IEnvironment) {
	shell.environment = &env
}

func (shell *Shell) GetEnvironment() sfinterfaces.IEnvironment {
	return *shell.environment
}

func (shell *Shell) SetSession(sess sfinterfaces.ISession) {
	shell.session = &sess
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
