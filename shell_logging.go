package shellframework

import (
	"fmt"
	"log"
	"os"
	"github.com/eshu0/shellframework/interfaces"
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

func (ssl *SimpleShellLog) SetLogPrefix(prefix string) {
	if !PointerInvalid(ssl.log) {
		ssl.log.SetPrefix(prefix)
	}
}

func (ssl *SimpleShellLog) SetLogLevel(lvl int) {
	ssl.loglevel = lvl
}

func (ssl *SimpleShellLog) GetLogLevel() int {
	return ssl.loglevel
}

func (ssl *SimpleShellLog) LogDebug(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogTrace(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogWarn(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogInfo(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}
func (ssl *SimpleShellLog) LogError(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogFatal(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogPrintln(msg string) {
	if !PointerInvalid(ssl.log) {
		ssl.log.Println(msg)
	}
}

func (ssl *SimpleShellLog) LogPrint(msg string) {
	if !PointerInvalid(ssl.log) {
		ssl.log.Print(msg)
	}
}

func (ssl *SimpleShellLog) LogPrintlnf(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogPrintf(msg string, a ...interface{}) {
	ssl.LogPrint(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) OpenSessionFileLog(session sfinterfaces.ISession) (*os.File) {
	f, err := os.OpenFile("simpleshell.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	log := log.New(f, session.ID()+" ", log.LstdFlags)

	// check log is valid
	if log == nil {
		panic("log is nil")
	}
	ssl.log = log
	ssl.loglevel = -1
	return f
}
