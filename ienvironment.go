package shellframework

type IEnvironment interface {
	GetShell() IShell
	SetShell(shell IShell)

	MakeSingleVariable(key string, val string) IEnvironmentVariable
	MakeMultiVariable(key string, vals []string) IEnvironmentVariable

	GetNameValues() map[string]IEnvironmentVariable
	SetNameValues(namval map[string]IEnvironmentVariable)

	SaveToFile(path string)
	LoadFile(path string)

	SetVariable(value IEnvironmentVariable)
	GetVariable(key string) (IEnvironmentVariable, bool)

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
