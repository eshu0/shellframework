package main

import (
	"os"
 	"flag"
	"github.com/eshu0/shellframework"
	"github.com/eshu0/simplelogger"
)

func main() {

	// grab the input
	scriptinfilepath := flag.String("in", "", "a file path to a script to be input")
	scriptoutfilepath := flag.String("out", "", "a file path to a script to be input")

	// parse the flags input to the tool
	flag.Parse()

	// create a new session (this will be a randomly generated id)
	session := shellframework.NewDefaultInteractiveSession(shellframework.DefaultBuildIDMethod)

	// lets open a flie log using the session
	logger := simplelogger.NewSimpleLogger("simpleshell.log", session.ID())

	// lets open a flie log using the session
	logger.OpenAllChannels()

	//defer the close till the shell has closed
	defer logger.CloseAllChannels()

	// default to standard in and standard out
	input :=  os.Stdin
	output := os.Stdout

	// script file been provided as input to the cli?
	if(*scriptoutfilepath != ""){

		// open/create the file if we cannot then crash out
		otf, err := os.OpenFile(*scriptoutfilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// should write an error here
			return
		}

		// replace the Stdin with the file we have read
		output = otf
	}

	var shell *shellframework.Shell

	if(*scriptinfilepath != ""){
		//Create a default shell using a file (in,out and err)
		shell = shellframework.NewShellFromFile(session,*scriptinfilepath, output, os.Stderr, &logger)
	}else{
		//Create a default shell using std (in,out and err)
		shell = shellframework.NewShell(session,input, output, os.Stderr, &logger)
	}

	// run the terminal
	// this would take input and read it
	// forever for loop
	shell.Run()

}
