# Shell Framework

Framework for a simple in-memory Golang Shell

This mainly is for small interactive CLI applications

In this example a shell is created that pipes to and from std in and out
Default commands can be found via help and a dummy session is created

```

package main

import (
	"github.com/eshu0/simpleshell"
	"github.com/eshu0/simpleshell/defaults"
)

func main() {

	//Create a default shell using std (in,out and err)
	// also log file is opened
	shell,f := ssdefaults.NewDefaultShellWithLog()

	// run the terminal
	// this would take input and read it forever
	shell.Run()

	//defer the close of the log file till the shell has closed
	defer f.Close()
}
```

In this example a shell is created that will execute the file input and will write the results to std out

```
package main

import (
	"os"

	"github.com/eshu0/simpleshell"
	"github.com/eshu0/simpleshell/defaults"
)

func main() {

	// create a session
	sss := simpleshell.NewSimpleSession()

	// create a log file, which will have the session id associated with it
	f1, log := ssdefaults.OpenIDBasedLog(sss)

	// open the file
	sh, err := os.Open("testrun.sh")
	if err != nil {
		return
	}

	// create a simple log
	ssl := simpleshell.NewSimpleShellLog(log)

	//Create a default shell and pass in the file
	shell := ssdefaults.NewDefaultShell(sss, sh, os.Stdout, os.Stderr, ssl)

	// run the terminal
	// this will execute the commands in the file
	shell.Run()

	//defer the close till the shell has closed
	defer f1.Close()

}

```
