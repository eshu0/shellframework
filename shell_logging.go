package shellframework

import (
	"fmt"
	"os"

	"github.com/eshu0/shellframework/interfaces"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type ShellLogger struct {
	sfinterfaces.IShellLogger
	loglevel kitlevel.Option
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

//func (ssl *ShellLogger) SetLogPrefix(prefix string) {
//ssl.log = kitlog.With(ssl.log, "session_id", session.ID())
//}

func (ssl *ShellLogger) SetLog(log kitlog.Logger) {
	ssl.log = log
}

func (ssl *ShellLogger) GetLog() kitlog.Logger {
	return ssl.log
}

func (ssl *ShellLogger) SetLogLevel(lvl kitlevel.Option) {
	ssl.loglevel = lvl
	// have to set the filter for the level
	ssl.log = kitlevel.NewFilter(ssl.log, lvl)
}

func (ssl *ShellLogger) GetLogLevel() kitlevel.Option {
	return ssl.loglevel
}

// the logging functions are here
func (ssl *ShellLogger) LogDebug(cmd string, data ...interface{}) {
	kitlevel.Debug(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

func (ssl *ShellLogger) LogWarn(cmd string, data ...interface{}) {
	kitlevel.Warn(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

func (ssl *ShellLogger) LogInfo(cmd string, data ...interface{}) {
	kitlevel.Info(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}
func (ssl *ShellLogger) LogError(cmd string, data ...interface{}) {
	kitlevel.Error(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

// the logging functions are here
func (ssl *ShellLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	kitlevel.Debug(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *ShellLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	kitlevel.Warn(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *ShellLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	kitlevel.Info(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}
func (ssl *ShellLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	kitlevel.Error(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *ShellLogger) OpenSessionFileLog(session sfinterfaces.ISession) *os.File {
	f, err := os.OpenFile("simpleshell.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	//logger :=
	logger := kitlog.NewLogfmtLogger(f)                                            //(f, session.ID()+" ", log.LstdFlags)
	logger = kitlog.With(logger, "session_id", session.ID(), "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

	// check log is valid
	if logger == nil {
		panic("logger is nil")
	}

	ssl.log = logger

	// default to show everything
	ssl.SetLogLevel(kitlevel.AllowAll())

	//ssl.loglevel = -1
	return f
}
