package shellframework

// main interface for the ShellLogger
type IShellLogger interface {

	//log functions
	GetLogLevel() int
	SetLogLevel(int)
	SetLogPrefix(string)

	LogFatal(string, ...interface{})
	LogError(string, ...interface{})
	LogWarn(string, ...interface{})
	LogInfo(string, ...interface{})
	LogDebug(string, ...interface{})
	LogTrace(string, ...interface{})

	LogPrintln(msg string)
	LogPrintlnf(msg string, a ...interface{})
	LogPrint(msg string)
	LogPrintf(msg string, a ...interface{})
}
