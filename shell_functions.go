package shellframework

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/eshu0/shellframework/interfaces"
)

//
// SHELL Printing
//
// these function provide printing to the Out
//

func (shell *Shell) Println(msg string) {
	if !PointerInvalid(shell.out) {
		shell.out.WriteString(msg + "\n")
	}
}

func (shell *Shell) Printlnf(msg string, a ...interface{}) {
	shell.Println(fmt.Sprintf(msg, a...))
}

func (shell *Shell) Printf(msg string, a ...interface{}) {
	shell.Print(fmt.Sprintf(msg, a...))
}

func (shell *Shell) Print(msg string) {
	if !PointerInvalid(shell.out) {
		shell.out.WriteString(msg)
	}
}

//
// SHELL Extra Print functions
//

func (shell *Shell) PrintDetails() {
	shell.Println("*****************************")
	shell.Printlnf("Framework Version: %s", shell.GetVersion())
	shell.Printlnf("Session: %s", shell.GetSession().ID())
	shell.Println("*****************************")
	shell.Println("")
}

func (shell *Shell) PrintInputMessage() {
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

func (shell *Shell) ParseInput(input string) []sfinterfaces.ICommandInput {

	var ecs []sfinterfaces.ICommandInput
	ecs = []sfinterfaces.ICommandInput{}
	log := *shell.GetLog()
	log.LogDebugf("ParseInput()", "Parsing '%s' with length %d", input, len(input))

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
			log.LogDebug("ParseInput()", "Comment found at the beggining this whole input is a comment finish parsing ")
			break
		} else {

			// first position let's create an command input
			if pos == 0 {
				ecsposition = 0
				commandfound = false
				commandfoundat = 0

				ci := CommandInput{}
				ecs = append(ecs, &ci)

				log.LogDebug("ParseInput()", "appended first command")
			}

			//shell.LogPrintlnf("character %c at position %d", char, pos)

			if char == '#' {
				log.LogDebugf("ParseInput()", "'%c' - Comment indentifier found at '%d' parsing finished", char, pos)
				break
			}

			// run out of string
			if len(input)-1 == pos {
				log.LogDebug("ParseInput()", "run out of string to parse")

				if !commandfound {
					cmndname := string(textr[commandfoundat : pos+1])
					log.LogDebugf("ParseInput()", "Parsed command '%s' from '%s'", cmndname, input)
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
					log.LogDebugf("ParseInput()", "Found argument terminator: argument read: %s", arg)
					pargs := ecs[ecsposition].GetArgs()
					pargs = append(pargs, arg)
					ecs[ecsposition].SetArgs(pargs)
				}

				rawi := string(textr[rawpos : pos+1])
				log.LogDebugf("ParseInput()", "Parsed rawinput '%s' from '%s'", rawi, input)
				ecs[ecsposition].SetRawInput(rawi)

				break

			} else {

				// we are looking for a command
				if !commandfound {
					if char == ' ' {
						log.LogDebugf("ParseInput()", "Final character %c at position %d is end of command", char, pos)
						cmndname := string(textr[commandfoundat : pos+1])
						log.LogDebugf("ParseInput()", "Parsed command '%s' from '%s'", cmndname, input)
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
							log.LogDebugf("ParseInput()", "Open qoute found at '%d' this is closing", pos)
							openqoute = false
						} else {
							log.LogDebugf("ParseInput()", "Open qoute found at '%d' this is opening", pos)
							openqoute = true
						}
						continue
					}

					// we keep going till it is closed
					if !openqoute {

						if char == ' ' {
							s := lastargat
							e := pos
							log.LogDebugf("ParseInput()", "e = %d", e)

							if textr[lastargat] == '"' {
								s = s + 1
							}
							log.LogDebugf("ParseInput()", "textr[e] = %s", string(textr[e]))

							if textr[e-1] == '"' {
								log.LogDebugf("ParseInput()", "minus e = %d", e)
								e = e - 1
							}

							log.LogDebugf("ParseInput()", "e = %d", e)
							arg := string(textr[s:e])
							log.LogDebugf("ParseInput()", "Found argument terminator: argument read: %s", arg)
							pargs := ecs[ecsposition].GetArgs()
							pargs = append(pargs, arg)
							ecs[ecsposition].SetArgs(pargs)
							lastargat = pos + 1
						}

						if char == '|' {
							log.LogDebugf("ParseInput()", "Pipe found at '%d' - new command input created", pos)
							log.LogDebugf("ParseInput()", "Append %s ", ecs[ecsposition].GetCommandName())
							ci := CommandInput{}
							ecs = append(ecs, &ci)
							ecsposition++
							commandfound = false

							rawi := string(textr[rawpos : pos+1])
							log.LogDebugf("ParseInput()", "Parsed rawinput '%s' from '%s'", rawi, input)
							ecs[ecsposition].SetRawInput(rawi)

							//this is zero on the first run so we need it to be past the pipe
							commandfoundat = pos + 1 // pos is the pipe command is at the next item
							lastargat = pos + 1
							rawpos = pos + 1
							log.LogDebugf("ParseInput()", "Created new command and incremented to %d ", ecsposition)

						}
					}

				}

			}

		}

	}

	log.LogDebug("ParseInput()", "following command input parsed and will be executed in order 0> ")
	for epos, cmdi := range ecs {
		log.LogDebugf("ParseInput()", "%d - Command: %s", epos, cmdi.GetCommandName())
		log.LogDebugf("ParseInput()", "%d - Raw Input: %s", epos, cmdi.GetRawInput())
		log.LogDebugf("ParseInput()", "%d - Input with out name: %s", epos, cmdi.GetInputWithOutCommand())

		for apos, arg := range cmdi.GetArgs() {
			log.LogDebugf("ParseInput()", "Args[%d]: %s", apos, arg)
		}
	}
	/*
		shell.LogPrintlnf("ParseInput(): Splitting '%s' by the pipe |", input)

			commands := strings.Split(input, "|")

			shell.LogPrintlnf("ParseInput(): Found '%d' commands ", len(commands))

			//args := strings.Split(text, " ")
			//shouldcontinue = cmd.Process(args[1:])

			for _, text := range commands {

				commentindex := strings.Index(text, "#")
				if commentindex > -1 {
					shell.LogPrintlnf("ParseInput(): Comment found at '%d' stripping after this ", commentindex)
					shell.LogPrintlnf("ParseInput(): string before parsing was %s ", text)
					textr := []rune(text)
					text = string(textr[:commentindex])
					shell.LogPrintlnf("ParseInput(): string after parsing was %s ", text)
				}

				// we have removed the comment
				// if the whole line was a comment we can ignore it
				if text == "" {
					shell.LogPrintln("ParseInput(): After removing comment line was empty - skipping ")
				} else {

					// filter out any silly caps lock mistakes
					lowerinput := strings.ToLower(text)

					// not sure this is the best thing to do
					// this could be made more comperhensive
					shell.LogPrintlnf("ParseInput(): Command '%s' matched '%s'", cmd.GetName(), lowercmd)

					runes := []rune(text)
					commandlength := len(cmd.GetName())

					shell.LogPrintlnf("ParseInput(): lowercmd length: %d ", commandlength)

					withoutcommand := string(runes[commandlength:])
					shell.LogPrintlnf("ParseInput(): without command %s", withoutcommand)
					withoutcommand = strings.TrimPrefix(withoutcommand, " ")
				}
			}
			/*
				if i > -1 {
					chars := x[:i]
					arefun := x[i+1:]
					fmt.Println(chars)
					fmt.Println(arefun)
				} else {
					fmt.Println("Index not found")
					fmt.Println(x)
				}
	*/
	return ecs
}

func (shell *Shell) Run() {
	// grab the environment
	env := shell.GetEnvironment()

	// pointer is valid?
	if !PointerInvalid(env) {
		env.LoadFile(sfinterfaces.EnvironmentFilename)
	}

	shell.PrintDetails()
	log := *shell.GetLog()
	shouldcontinue := true
	session := shell.GetSession()

	if session.GetInteractive() {
		log.LogDebug("Run()", "Interactive Session")
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
				char, key, err := keyboard.GetSingleKey() // keyboard.GetKey()
				if err != nil {
					panic(err)
				} else if key == keyboard.KeyArrowUp {
					if !PointerInvalid(env) {
						envvar, exists := env.GetVariable(sfinterfaces.LastCommands)
						if exists {
							wc := envvar
							lc := wc.GetValues()
							if lastcommandpos >= len(lc)-1 {
								shell.PrintInputMessage()
								shell.Printf("%s", lc[lastcommandpos])
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
						envvar, exists := env.GetVariable(sfinterfaces.LastCommands)
						if exists {
							wc := envvar
							lc := wc.GetValues()
							if lastcommandpos >= len(lc)-1 {
								shell.PrintInputMessage()
								shell.Printf("%s", lc[lastcommandpos])
								lastcommandpos = 0
							} else {
								shell.PrintInputMessage()
								shell.Printf("%s", lc[lastcommandpos])
								lastcommandpos = lastcommandpos - 1
							}
						}

					}
				} else if key == keyboard.KeyEsc {
					shell.Println("Exiting")
					log.LogDebug("Run()", "Exiting")
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

				env.AddStringValue(sfinterfaces.LastCommands, text)
				/*
					envvar, exists := env.GetVariable(sfinterfaces.LastCommands)

					if !exists {
						var cmds []string
						cmds = append(cmds, text)
						env.SetVariable(env.MakeMultiVariable(sfinterfaces.LastCommands, cmds))
					} else {
						wc := envvar
						lc := wc.GetValues()
						lc = append(lc, text)
						wc.SetValues(lc)
						env.SetVariable(wc)
					}
				*/

			}

			executionorder := shell.ParseInput(text)

			var cmdres string
			endexecution := false
			for _, ec := range executionorder {

				log.LogDebugf("Run()", "Found '%s' execution command  ", ec.GetCommandName())
				if endexecution {
					break
				}

				if cmdres != "" {
					log.LogDebugf("Run()", "Previous command finished with result %s override the args", cmdres)
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
						log.LogDebugf("Run()", "Started SetCommandInput for '%s' ", cmd.GetName())
						cmd.SetCommandInput(ec)
						log.LogDebugf("Run()", "Finished SetCommandInput for '%s' ", cmd.GetName())

						log.LogDebugf("Run()", "Started command '%s' ", cmd.GetName())
						res := cmd.Process()
						log.LogDebugf("Run()", "Finished command '%s'  ", cmd.GetName())

						if res.ExitShell() {
							shouldcontinue = false
						} else {

							if res.Sucessful() {
								log.LogDebugf("Run()", "Command '%s' was sucessful ", cmd.GetName())
								cmdres = res.Result()
							} else {
								err := res.Err()
								if err != nil {
									shell.Printlnf("'%s' failed: %s ", cmd.GetName(), err.Error())
									log.LogDebugf("Run()", "Error with command '%s' following error provided: %s ", cmd.GetName(), err.Error())
								} else {
									shell.Printlnf("Error with command '%s' no error provided ", cmd.GetName())
									log.LogDebugf("Run()", "Error with command '%s' no error provided  ", cmd.GetName())
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
				log.LogDebug("Run()", "Exiting")
				break
			}

			if !PointerInvalid(env) {
				env.SaveToFile(sfinterfaces.EnvironmentFilename)
			}

		} // for loop

	} else {
		log.LogDebug("Run()", "Non-Interactive Session")

		reader := bufio.NewReader(shell.in)
		for {

			// this keeps updating so let's keep it syncd
			env = shell.GetEnvironment()

			// pointer is valid?
			if !PointerInvalid(reader) {

				// print the input message
				shell.PrintInputMessage()

				// read the string input
				text, readerr := reader.ReadString('\n')

				if readerr != nil {
					log.LogDebugf("Run()", "Reading input has provided following err '%s'", readerr.Error())
					break
					// break out for loop
				}

				// convert CRLF to LF
				text = strings.Replace(text, "\n", "", -1)

				// pointer is valid?
				if !PointerInvalid(env) && text != "" && (len(text) > 0 && text[0] != '#') {

					env.AddStringValue(sfinterfaces.LastCommands, text)
					/*
						envvar, exists := env.GetVariable(sfinterfaces.LastCommands)

						if !exists {
							var cmds []string
							cmds = append(cmds, text)
							env.SetVariable(env.MakeMultiVariable(sfinterfaces.LastCommands, cmds))
						} else {
							wc := envvar
							lc := wc.GetValues()
							lc = append(lc, text)
							wc.SetValues(lc)
							env.SetVariable(wc)
						}
					*/
				}

				executionorder := shell.ParseInput(text)

				var cmdres string
				endexecution := false
				for _, ec := range executionorder {

					log.LogDebugf("Run()", "Found '%s' execution command  ", ec.GetCommandName())
					if endexecution {
						break
					}

					if cmdres != "" {
						log.LogDebugf("Run()", "Previous command finished with result %s override the args", cmdres)
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
							cmd.SetCommandInput(ec)

							res := cmd.Process()

							if res.ExitShell() {
								shouldcontinue = false
							} else {

								if res.Sucessful() {
									cmdres = res.Result()
								} else {
									err := res.Err()
									if err != nil {
										shell.Printlnf("'%s' failed: %s ", cmd.GetName(), err.Error())
										log.LogDebugf("Run()", "Error with command '%s' following error provided: %s ", cmd.GetName(), err.Error())
									} else {
										shell.Printlnf("Error with command '%s' no error provided ", cmd.GetName())
										log.LogDebugf("Run()", "Error with command '%s' no error provided  ", cmd.GetName())
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
					log.LogDebug("Run()", "Exiting")
					break
				}

			} else {

				log.LogDebug("Run()", "Reader is nil")
				shouldcontinue = false
			}

			if !PointerInvalid(env) {
				env.SaveToFile(sfinterfaces.EnvironmentFilename)
			}

		} // for loop
	}

}

//
// Commands adding etc
//

func (shell *Shell) NewCommand(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult, flags []sfinterfaces.IFlag) sfinterfaces.ICommand {

	sc := &Command{}
	sc.name = name
	sc.operator = operator //
	sc.description = description
	sc.shell = shell

	flgs := &CommandFlags{}
	flgs.SetCommand(sc)
	flgs.SetFlags(flags)

	sc.SetCommandFlags(flgs)

	//sc.flags = flags
	return sc
}

func (shell *Shell) AddCommand(cmd sfinterfaces.ICommand) {
	// append the command to the shell
	shell.commands = append(shell.commands, cmd)
}

//Adds a Simple Command to the Shell
func (shell *Shell) AddCommands(commands []sfinterfaces.ICommand) {
	// walk thoguh the commands passed in
	for _, cmd := range commands {
		// make sure this pointer is valid
		if !PointerInvalid(cmd) {
			// use this method to add the simple command
			shell.AddCommand(cmd)
		}
	}

}

func (shell *Shell) RegisterNewCommandWithFlags(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult, flags []sfinterfaces.IFlag) {
	shell.AddCommand(shell.NewCommand(name, description, operator, flags))
}

func (shell *Shell) RegisterNewCommand(name string, description string, operator func(command sfinterfaces.ICommand) sfinterfaces.ICommandResult) {
	flags := []sfinterfaces.IFlag{}
	shell.AddCommand(shell.NewCommand(name, description, operator, flags))
}

func (shell *Shell) RegisterCommandNewBoolFlag(cmd string, name string, defaultvalue bool, usage string) {
	sf := &CommandFlag{}
	sf.name = name
	sf.defaultbvalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 2

	shell.RegisterCommandFlag(cmd, sf)
}

func (shell *Shell) RegisterCommandNewIntFlag(cmd string, name string, defaultvalue int, usage string) {
	sf := &CommandFlag{}
	sf.name = name
	sf.defaultivalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 3

	shell.RegisterCommandFlag(cmd, sf)
}

func (shell *Shell) RegisterCommandNewStringFlag(cmd string, name string, defaultvalue string, usage string) {

	sf := &CommandFlag{}
	sf.name = name
	sf.defaultsvalue = defaultvalue
	sf.usage = usage
	sf.flagtype = 1
	shell.RegisterCommandFlag(cmd, sf)
}

func (shell *Shell) RegisterCommandFlag(cmd string, flag sfinterfaces.IFlag) {

	log := *shell.GetLog()

	for i, _ := range shell.commands {
		// make sure this pointer is valid

		if strings.ToLower(shell.commands[i].GetName()) == strings.ToLower(cmd) {
			log.LogDebugf("RegisterCommandFlag()", "'%s' macthed '%s'", shell.commands[i].GetName(), cmd)

			flgsbefore := shell.commands[i].GetCommandFlags().GetFlags()
			for p, flg := range flgsbefore {
				log.LogDebugf("RegisterCommandFlag()", "flgsbefore - commands[%d][%d] has '%s' set to type %d", i, p, flg.GetName(), flg.GetFlagType())
			}

			// get the iflags from the command
			flgs := shell.commands[i].GetCommandFlags()

			// now get the underlying array list
			flags := flgs.GetFlags()
			flags = append(flags, flag)
			flgs.SetFlags(flags)

			log.LogDebugf("RegisterCommandFlag()", "commands[%d] had it's flags set", i)
			shell.commands[i].SetCommandFlags(flgs)

			flgsafter := shell.commands[i].GetCommandFlags().GetFlags()
			for j, flg := range flgsafter {
				log.LogDebugf("RegisterCommandFlag()", "flgsafter - commands[%d][%d] has '%s' set to type %d", i, j, flg.GetName(), flg.GetFlagType())
			}
			return

		}
	}
}
