package sfinterfaces

// the session for the Shell
type ISession interface {
	//The Identifying String
	ID() string

	// Session is Interactive?
	GetInteractive() bool
	SetInteractive(interactive bool)

}
