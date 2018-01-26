package main

import (
	"fmt"

	"github.com/rubencaro/omg/lib/cli"
	"github.com/rubencaro/omg/lib/input"
)

var version string // to be set from build script

func main() {
	// Read data
	data, err := input.Read()
	if err != nil {
		fmt.Println("We need configuration! \n", err)
		return
	}

	data.Version = version

	// Start CLI
	err = cli.Execute(data)
	if err != nil {
		fmt.Println("OMG", err)
	}
}
