package simpleshell

import (
	"fmt"
	"log"
)

type SimpleShellLog struct {
	loglevel int
	log      *log.Logger
}

//
// SHELL Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleShellLog(logger *log.Logger) SimpleShellLog {
	ssl := SimpleShellLog{}
	ssl.log = logger
	ssl.loglevel = -1
	return ssl
}

func (shell *SimpleShellLog) SetLogPrefix(prefix string) {
	if !PointerInvalid(shell.log) {
		shell.log.SetPrefix(prefix)
	}
}

func (shell *SimpleShellLog) SetLogLevel(lvl int) {
	shell.loglevel = lvl
}

func (shell *SimpleShellLog) GetLogLevel() int {
	return shell.loglevel
}

func (shell *SimpleShellLog) LogDebug(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShellLog) LogTrace(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShellLog) LogWarn(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShellLog) LogInfo(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}
func (shell *SimpleShellLog) LogError(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))

}

func (shell *SimpleShellLog) LogFatal(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShellLog) LogPrintln(msg string) {
	if !PointerInvalid(shell.log) {
		shell.log.Println(msg)
	}
}

func (shell *SimpleShellLog) LogPrint(msg string) {
	if !PointerInvalid(shell.log) {
		shell.log.Print(msg)
	}
}

func (shell *SimpleShellLog) LogPrintlnf(msg string, a ...interface{}) {
	shell.LogPrintln(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShellLog) LogPrintf(msg string, a ...interface{}) {
	shell.LogPrint(fmt.Sprintf(msg, a...))
}
