# Simple shell

Simple in-memory Golang Shell

This mainly is for small interactive CLI applications

In this example a shell is created that pipes to and from std in and out
Default commands can be found via help and a dummy session is created

```
package main

import (
	"os"
 	"flag"
	"github.com/eshu0/shellframework"
)

func main() {

	// grab the input
	scriptinfilepath := flag.String("in", "", "a file path to a script to be input")
	scriptoutfilepath := flag.String("out", "", "a file path to a script to be input")

	// parse the flags input to the tool
	flag.Parse()

	// create a new session (this will be a randomly generated id)
	session := shellframework.NewSession()

	// this is the dummy logger object
	logger := shellframework.ShellLogger{}

	// lets open a flie log using the session
	f1 := logger.OpenSessionFileLog(session)

	// default to standard in and standard out
	input :=  os.Stdin
	output := os.Stdout

	// script file been provided as input to the cli?
	if(*scriptinfilepath != ""){

		// open the file if we cannot then crash out
		inf, err := os.Open(*scriptinfilepath)
		if err != nil {
			// should write an error here
			return
		}

		// replace the Stdin with the file we have read
		input = inf
	}

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

	//Create a default shell using std (in,out and err)
	shell := shellframework.NewShell(session,input, output, os.Stderr, &logger)

	// run the terminal
	// this would take input and read it
	// forever for loop
	shell.Run()

	//defer the close till the shell has closed
	defer f1.Close()
}

```
