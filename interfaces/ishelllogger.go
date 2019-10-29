package sfinterfaces

import (
	"os"

	kitlog "github.com/go-kit/kit/log"
		kitlevel "github.com/go-kit/kit/log/level"
)

// main interface for the ShellLogger
type IShellLogger interface {
	// logging
	GetLog() kitlog.Logger
	SetLog(log kitlog.Logger)

	//log functions
	GetLogLevel() kitlevel.Option
	SetLogLevel(kitlevel.Option)
	//SetLogPrefix(string)

	LogError(data ...interface{})
	LogWarn(data ...interface{})
	LogInfo(data ...interface{})
	LogDebug(data ...interface{})

	LogPrintln(msg string)
	LogPrintlnf(msg string, a ...interface{})
	LogPrint(msg string)
	LogPrintf(msg string, a ...interface{})

	OpenSessionFileLog(session ISession) *os.File
}
