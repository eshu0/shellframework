package shellframework

import (
	"fmt"
	"os"

	"github.com/eshu0/shellframework/interfaces"
	kitlog "github.com/go-kit/kit/log"
)

type ShellLogger struct {
	loglevel int
	log      kitlog.Logger
}

//
// SHELL Logging
//
// these function provide logging to the choosen logfile
//

func NewShellLogger(logger kitlog.Logger) ShellLogger {
	ssl := ShellLogger{}
	ssl.log = logger
	return ssl
}

func (ssl *ShellLogger) SetLogPrefix(prefix string) {
	//ssl.log = kitlog.With(ssl.log, "session_id", session.ID())
}

func (ssl *ShellLogger) SetLog(log kitlog.Logger) {
	ssl.log = log
}

func (ssl *ShellLogger) GetLog() kitlog.Logger {
	return ssl.log
}

func (ssl *ShellLogger) SetLogLevel(lvl int) {
	ssl.loglevel = lvl
}

func (ssl *ShellLogger) GetLogLevel() int {
	return ssl.loglevel
}

func (ssl *ShellLogger) LogDebug(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogTrace(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogWarn(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogInfo(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}
func (ssl *ShellLogger) LogError(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogFatal(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogPrintln(msg string) {
	ssl.log.Log(fmt.Sprintf("%s \n", msg))
}

func (ssl *ShellLogger) LogPrint(msg string) {
	ssl.log.Log(msg)
}

func (ssl *ShellLogger) LogPrintlnf(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) LogPrintf(msg string, a ...interface{}) {
	ssl.LogPrint(fmt.Sprintf(msg, a...))
}

func (ssl *ShellLogger) OpenSessionFileLog(session sfinterfaces.ISession) *os.File {
	f, err := os.OpenFile("simpleshell.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	//logger :=
	logger := kitlog.NewLogfmtLogger(f) //(f, session.ID()+" ", log.LstdFlags)
	logger = kitlog.With(logger, "session_id", session.ID())

	// check log is valid
	if logger == nil {
		panic("logger is nil")
	}
	ssl.log = logger
	ssl.loglevel = -1
	return f
}
