package sfinterfaces

// the session for the Shell
type ISession interface {
	//The Identifying String
	SetIDMethod(idmethod func() string)
	GetIDMethod() func() string

	// Session is Interactive?
	SetInteractiveMethod(interactivemethod func() bool)
	GetInteractiveMethod() func() bool

	//Get and Set for the shell that this command belong to
	SetShell(shell IShell)
	GetShell() IShell

}
