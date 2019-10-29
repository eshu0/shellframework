package shellframework

import (
	"os"

	"github.com/eshu0/shellframework/defaultcmds"
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

	// Add the default commands
	//shell.AddNewCommand("man", "Manual command similar to linux/unix systems", dcmds.Man)
	cmd := dcmds.ManCommand{}
	cmd.Register(shell)

	shell.AddNewCommand("help", "Help command", dcmds.Help)
	shell.AddNewCommand("exit", "Exit terminal command", dcmds.Exit)
	shell.AddNewCommand("cmd", "Commands command - jinx!", dcmds.Cmd)

	Flags := []sfinterfaces.IFlag{}
	//flg :=
	Flags = append(Flags, NewStringFlag("key", "", "Sets a string value"))
	Flags = append(Flags, NewStringFlag("value", "", "Sets a string value"))
	Flags = append(Flags, NewBoolFlag("list", false, "List Environment Variables"))

	shell.AddNewCommandWithFlags("env", "Environment command", dcmds.Env, Flags)
	shell.AddNewCommand("echo", "Echo text to terminal", dcmds.Echo)

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
