package shellframework

import (
	"fmt"
	"os"

	"github.com/eshu0/shellframework/interfaces"
	kitlog "github.com/go-kit/kit/log"
)

type SimpleShellLog struct {
	loglevel int
	log      kitlog.Logger
}

//
// SHELL Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleShellLog(logger kitlog.Logger) SimpleShellLog {
	ssl := SimpleShellLog{}
	ssl.log = logger
	return ssl
}

func (ssl *SimpleShellLog) SetLogPrefix(prefix string) {
	//ssl.log = kitlog.With(ssl.log, "session_id", session.ID())
}

func (ssl *SimpleShellLog) SetLog(log kitlog.Logger) {
	ssl.log = log
}

func (ssl *SimpleShellLog) GetLog() kitlog.Logger {
	return ssl.log
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
	ssl.log.Log(fmt.Sprintf("%s \n", msg))
}

func (ssl *SimpleShellLog) LogPrint(msg string) {
	ssl.log.Log(msg)
}

func (ssl *SimpleShellLog) LogPrintlnf(msg string, a ...interface{}) {
	ssl.LogPrintln(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) LogPrintf(msg string, a ...interface{}) {
	ssl.LogPrint(fmt.Sprintf(msg, a...))
}

func (ssl *SimpleShellLog) OpenSessionFileLog(session sfinterfaces.ISession) *os.File {
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
