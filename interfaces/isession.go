package sfinterfaces

// the session for the Shell
type ISession interface {
	//The Identifying String
	SetIDMethod(idmethod func(ss ISession) string)
	GetIDMethod() func(ss ISession) string

	// Session is Interactive?
	SetInteractiveMethod(interactivemethod func(ss ISession) bool)
	GetInteractiveMethod() func(ss ISession) bool

	//Get and Set for the shell that this command belong to
	SetShell(shell IShell)
	GetShell() IShell

}
