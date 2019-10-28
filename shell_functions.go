package shellframework

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/eshu0/shellframework/interfaces"
)

/*
*
*	DEFAULT SHELL ENVIRONMENT VARIABLES
*
 */
const EnvironmentFilename string = "env"
const LastCommands string = "LastCommands"
const PersistEnvironment string = "PersistEnvironment"

//
// SHELL Printing
//
// these function provide printing to the Out
//

func (shell *SimpleShell) Println(msg string) {
	if !PointerInvalid(shell.out) {
		shell.out.WriteString(msg + "\n")
	}
}

func (shell *SimpleShell) Printlnf(msg string, a ...interface{}) {
	shell.Println(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShell) Printf(msg string, a ...interface{}) {
	shell.Print(fmt.Sprintf(msg, a...))
}

func (shell *SimpleShell) Print(msg string) {
	if !PointerInvalid(shell.out) {
		shell.out.WriteString(msg)
	}
}

//
// SHELL Extra Print functions
//

func (shell *SimpleShell) PrintDetails() {
	shell.Println("*****************************")
	shell.Printlnf("Version: %s", shell.GetVersion())
	shell.Printlnf("Session: %s", shell.GetSession().ID())
	shell.Println("*****************************")
	shell.Println("")
}

func (shell *SimpleShell) PrintInputMessage() {
	sess := shell.GetSession()
	if !PointerInvalid(sess) {
		shell.Printf("[%s]: ", sess.ID())
	} else {
		shell.Print("[Invalid session]: ")
	}
}

//
// SHELL Processing
//

func (shell *SimpleShell) ParseInput(input string) []sfinterfaces.ICommandInput {

	var ecs []sfinterfaces.ICommandInput
	ecs = []sfinterfaces.ICommandInput{}
	log := *shell.GetLog()
	log.LogPrintlnf("ParseInput(): Parsing '%s' with length %d", input, len(input))

	var ecsposition int
	var commandfound bool
	var commandfoundat int

	//var pargs []string
	commandfound = false
	ecsposition = 0
	commandfoundat = 0

	var openqoute bool
	var lastargat int
	//var argstart int
	//var argend int

	openqoute = false
	lastargat = 0
	//argstart = 0
	//argend = 0

	var rawpos int
	rawpos = 0

	textr := []rune(input)
	for pos, char := range textr {

		if pos == 0 && char == '#' {
			log.LogPrintln("ParseInput():Comment found at the beggining this whole input is a comment finish parsing ")
			break
		} else {

			// first position let's create an command input
			if pos == 0 {
				ecsposition = 0
				commandfound = false
				commandfoundat = 0

				ci := SimpleCommandInput{}
				ecs = append(ecs, &ci)

				log.LogPrintln("ParseInput(): appended first command")
			}

			//shell.LogPrintlnf("character %c at position %d", char, pos)

			if char == '#' {
				log.LogPrintlnf("ParseInput(): '%c' - Comment indentifier found at '%d' parsing finished", char, pos)
				break
			}

			// run out of string
			if len(input)-1 == pos {
				log.LogPrintlnf("run out of string to parse")

				if !commandfound {
					cmndname := string(textr[commandfoundat : pos+1])
					log.LogPrintlnf("Parsed command '%s' from '%s'", cmndname, input)
					ecs[ecsposition].SetCommandName(cmndname)

				} else {
					s := lastargat
					e := pos

					if textr[lastargat] == '"' {
						s = s + 1
					}

					if textr[pos] == '"' {
						e = e - 1
					}

					arg := string(textr[s : e+1])
					log.LogPrintlnf("ParseInput(): Found argument terminator: argument read: %s", arg)
					pargs := ecs[ecsposition].GetArgs()
					pargs = append(pargs, arg)
					ecs[ecsposition].SetArgs(pargs)
				}

				rawi := string(textr[rawpos : pos+1])
				log.LogPrintlnf("Parsed rawinput '%s' from '%s'", rawi, input)
				ecs[ecsposition].SetRawInput(rawi)

				break

			} else {

				// we are looking for a command
				if !commandfound {
					if char == ' ' {
						log.LogPrintlnf("Final character %c at position %d is end of command", char, pos)
						cmndname := string(textr[commandfoundat : pos+1])
						log.LogPrintlnf("Parsed command '%s' from '%s'", cmndname, input)
						ecs[ecsposition].SetCommandName(cmndname)
						commandfoundat = pos
						commandfound = true
						lastargat = pos + 1
					} else {
						// let's keep looking
						continue
					}
				} else { // we are parsing arguements

					if char == '"' {
						if openqoute {
							log.LogPrintlnf("ParseInput(): Open qoute found at '%d' this is closing", pos)
							openqoute = false
						} else {
							log.LogPrintlnf("ParseInput(): Open qoute found at '%d' this is opening", pos)
							openqoute = true
						}
						continue
					}

					// we keep going till it is closed
					if !openqoute {

						if char == ' ' {
							s := lastargat
							e := pos
							log.LogPrintlnf("ParseInput(): e = %d", e)

							if textr[lastargat] == '"' {
								s = s + 1
							}
							log.LogPrintlnf("ParseInput(): textr[e] = %s", string(textr[e]))

							if textr[e-1] == '"' {
								log.LogPrintlnf("ParseInput(): minus e = %d", e)
								e = e - 1
							}

							log.LogPrintlnf("ParseInput(): e = %d", e)
							arg := string(textr[s:e])
							log.LogPrintlnf("ParseInput(): Found argument terminator: argument read: %s", arg)
							pargs := ecs[ecsposition].GetArgs()
							pargs = append(pargs, arg)
							ecs[ecsposition].SetArgs(pargs)
							lastargat = pos + 1
						}

						if char == '|' {
							log.LogPrintlnf("ParseInput(): Pipe found at '%d' - new command input created", pos)
							log.LogPrintlnf("ParseInput(): Append %s ", ecs[ecsposition].GetCommandName())
							ci := SimpleCommandInput{}
							ecs = append(ecs, &ci)
							ecsposition++
							commandfound = false

							rawi := string(textr[rawpos : pos+1])
							log.LogPrintlnf("Parsed rawinput '%s' from '%s'", rawi, input)
							ecs[ecsposition].SetRawInput(rawi)

							//this is zero on the first run so we need it to be past the pipe
							commandfoundat = pos + 1 // pos is the pipe command is at the next item
							lastargat = pos + 1
							rawpos = pos + 1
							log.LogPrintlnf("ParseInput(): Created new command and incremented to %d ", ecsposition)

						}
					}

				}

			}

		}

	}

	log.LogPrintln("ParseInput(): following command input parsed and will be executed in order 0> ")
	for epos, cmdi := range ecs {
		log.LogPrintlnf("ParseInput():  %d - Command: %s", epos, cmdi.GetCommandName())
		log.LogPrintlnf("ParseInput():  %d - Raw Input: %s", epos, cmdi.GetRawInput())
		log.LogPrintlnf("ParseInput():  %d - Input with out name: %s", epos, cmdi.GetInputWithOutCommand())

		for apos, arg := range cmdi.GetArgs() {
			log.LogPrintlnf("ParseInput():  Args[%d]: %s", apos, arg)
		}
	}

	return ecs
}

func (shell *SimpleShell) Run() {

	// grab the environment
	env := shell.GetEnvironment()

	// pointer is valid?
	if !PointerInvalid(env) {
		env.LoadFile(EnvironmentFilename)
	}

	shell.PrintDetails()
	log := *shell.GetLog()

	shouldcontinue := true
	lastcommandpos := 0
	for {

		// this keeps updating so let's keep it syncd
		env = shell.GetEnvironment()

		// print the input message
		shell.PrintInputMessage()

		kerr := keyboard.Open()
		if kerr != nil {
			panic(kerr)
		}
		defer keyboard.Close()
		text := ""
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			} else if key == keyboard.KeyArrowUp {
				if !PointerInvalid(env) {
					envvar, exists := env.GetVariable(LastCommands)
					if exists {
						wc := envvar
						lc := wc.GetValues()
						if lastcommandpos >= len(lc)-1 {
							shell.PrintInputMessage()
							shell.Printf(" %s", lc[lastcommandpos])
							lastcommandpos = 0
						} else {
							shell.PrintInputMessage()
							shell.Printf(" %s", lc[lastcommandpos])
							lastcommandpos = lastcommandpos + 1
						}
					}
				}
			} else if key == keyboard.KeyArrowDown {
				if !PointerInvalid(env) {
					envvar, exists := env.GetVariable(LastCommands)
					if exists {
						wc := envvar
						lc := wc.GetValues()
						if lastcommandpos >= len(lc)-1 {
							shell.PrintInputMessage()
							shell.Printf(" %s", lc[lastcommandpos])
							lastcommandpos = 0
						} else {
							shell.PrintInputMessage()
							shell.Printf(" %s", lc[lastcommandpos])
							lastcommandpos = lastcommandpos - 1
						}
					}

				}
			} else if key == keyboard.KeyEsc {
				shell.Println("Exiting")
				log.LogPrintln("Run(): Exiting")
				return
			} else if key == keyboard.KeyEnter {
				shell.Print("\n")
				break
			} else {
				shell.Print(string(char))
				text = text + string(char)
			}

		}

		// pointer is valid?
		if !PointerInvalid(env) {

			envvar, exists := env.GetVariable(LastCommands)

			if !exists {
				var cmds []string
				cmds = append(cmds, text)
				env.SetVariable(env.MakeMultiVariable(LastCommands, cmds))
			} else {
				wc := envvar
				lc := wc.GetValues()
				lc = append(lc, text)
				wc.SetValues(lc)
				env.SetVariable(wc)
			}

		}

		executionorder := shell.ParseInput(text)

		var cmdres string
		endexecution := false
		for _, ec := range executionorder {

			log.LogPrintlnf("Run(): Found '%s' execution command  ", ec.GetCommandName())
			if endexecution {
				break
			}

			if cmdres != "" {
				log.LogPrintlnf("Run(): Previous command finished with result %s override the args", cmdres)
				var pargs []string
				pargs = append(pargs, cmdres)
				ec.SetArgs(pargs)
			}

			// walk commands in shell
			for _, cmd := range shell.GetCommands() {
				// This command matched
				if cmd.Match(ec) {
					// not sure this is the best thing to do
					// this could be made more comperhensive
					// we set this here so prasing doesn;t affect the input
					log.LogPrintlnf("Run(): Started SetCommandInput for '%s' ", cmd.GetName())
					cmd.SetCommandInput(ec)
					log.LogPrintlnf("Run(): Finished SetCommandInput for '%s' ", cmd.GetName())

					log.LogPrintlnf("Run(): Started command '%s' ", cmd.GetName())
					res := cmd.Process()
					log.LogPrintlnf("Run(): Finished command '%s'  ", cmd.GetName())

					if res.ExitShell() {
						shouldcontinue = false
					} else {

						if res.Sucessful() {
							log.LogPrintlnf("Run(): Command '%s' was sucessful ", cmd.GetName())
							cmdres = res.Result()
						} else {
							err := res.Err()
							if err != nil {
								shell.Printlnf("'%s' failed: %s ", cmd.GetName(), err.Error())
								log.LogPrintlnf("Run(): Error with command '%s' following error provided: %s ", cmd.GetName(), err.Error())
							} else {
								shell.Printlnf("Error with command '%s' no error provided ", cmd.GetName())
								log.LogPrintlnf("Run(): Error with command '%s' no error provided  ", cmd.GetName())
							}
							endexecution = true
						}

					}

					break
				}
			}
		}

		if !shouldcontinue {
			shell.Println("Exiting")
			log.LogPrintln("Run(): Exiting")
			break
		}

		if !PointerInvalid(env) {
			env.SaveToFile(EnvironmentFilename)
		}

	} // for loop

}

//
// Commands adding etc
//
func (shell *SimpleShell) AddCommand(cmd sfinterfaces.ICommand) {
	// append the command to the shell
	shell.commands = append(shell.commands, cmd)
}

//Adds a Simple Command to the Shell
func (shell *SimpleShell) AddCommands(commands []sfinterfaces.ICommand) {
	// walk thoguh the commands passed in
	for _, cmd := range commands {
		// make sure this pointer is valid
		if !PointerInvalid(cmd) {
			// use this method to add the simple command
			shell.AddCommand(cmd)
		}
	}

}

func (shell *SimpleShell) AddNewCommandWithFlags(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult, flags []sfinterfaces.IFlag) {
	shell.AddCommand(shell.NewCommand(name, description, operator, flags))
}

func (shell *SimpleShell) AddNewCommand(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult) {
	flags := []sfinterfaces.IFlag{}
	shell.AddCommand(shell.NewCommand(name, description, operator, flags))
}

func (shell *SimpleShell) NewCommand(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult, flags []sfinterfaces.IFlag) sfinterfaces.ICommand {

	sc := &SimpleCommand{}
	sc.name = name
	sc.operator = operator //
	sc.description = description
	sc.shell = shell

	flgs := &SimpleFlags{}
	flgs.SetCommand(sc)
	flgs.SetFlags(flags)

	sc.SetFlags(flgs)

	//sc.flags = flags
	return sc
}
