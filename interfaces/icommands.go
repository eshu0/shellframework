package sfinterfaces

type ICommand interface {

	// This is the description of the commands
	GetDescription() string

	// This is the name of the command to be matched at command line
	GetName() string

	//Get and Set for the shell that this command belong to
	SetShell(shell IShell)
	GetShell() IShell

	//Get and Set IExecution
	SetCommandInput(input ICommandInput)
	GetCommandInput() ICommandInput

	// The command arguments that are required for parsing for this command
	//SetArguments(args ICommandArguments)
	//GetArguments() ICommandArguments

	// Match for matching the text input
	// it will be lower so that it is case insensitive
	Match(incmd ICommandInput) bool

	// if matched then we shall process the rest of the commands
	// this return a false if we should Exit
	// aka finish the shell
	// this can be made more complicated however bool for the first revision
	Process() ICommandResult

	//Get and Set IFlags
	SetFlags(flgs IFlags)
	GetFlags() IFlags

	// helper methods
	NewSuccessCommandResult(msg string) ICommandResult
	NewExitSuccessCommandResult() ICommandResult
}

// This represents the input from the command line
// it is important if a command wants to ignore the shells parsing
type ICommandInput interface {
	GetCommandName() string
	GetLowerCommandName() string
	GetInputWithOutCommand() string
	GetArgs() []string
	GetRawInput() string

	SetCommandName(name string)
	SetRawInput(name string)
	SetArgs(args []string)
}

// this represents the result of executing the command
type ICommandResult interface {
	ExitShell() bool
	Sucessful() bool
	Err() error
	Result() string
}
