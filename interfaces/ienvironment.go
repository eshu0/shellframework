package sfinterfaces

/*
*
*	DEFAULT SHELL ENVIRONMENT VARIABLES
*
 */
const EnvironmentFilename string = "env"
const PersistEnvironment string = "PersistEnvironment"
const LastCommands string = "LastCommands"

type IEnvironment interface {
	GetShell() IShell
	SetShell(shell IShell)

	MakeSingleVariable(key string, val string) IEnvironmentVariable
	MakeMultiVariable(key string, vals []string) IEnvironmentVariable

	GetNameValues() map[string]IEnvironmentVariable
	SetNameValues(namval map[string]IEnvironmentVariable)

	SaveToFile(path string)
	LoadFile(path string)

	Set(value IEnvironmentVariable)
	Clear(key string)
	Delete(key string)

	GetVariable(key string) (IEnvironmentVariable, bool)

	AddStringValue(key string, value string)

	Print()
}

type IEnvironmentVariable interface {
	GetName() string
	GetValues() []string
	GetType() int

	SetName(name string)
	SetValues(vals []string)
	SetType(typ int)
}
