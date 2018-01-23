package main

import (
	"fmt"

	"github.com/rubencaro/omg/lib/cmd"
	"github.com/rubencaro/omg/lib/cnf"
)

var version string // to be set from build script

func main() {
	// Read config
	conf, err := cnf.Read()
	if err != nil {
		fmt.Println("We need configuration! \n", err)
		return
	}

	// Start Cobra CLI
	err = cmd.Execute(conf, version)
	if err != nil {
		fmt.Println("OMG it failed! \n", err)
	}
}
