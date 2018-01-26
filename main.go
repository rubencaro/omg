package main

import (
	"fmt"

	"github.com/rubencaro/omg/lib/cli"
	"github.com/rubencaro/omg/lib/input"
)

var release = "0.1.0"
var commit string // to be set from build script

func main() {
	// Read data
	data, err := input.Read()
	if err != nil {
		fmt.Println("We need configuration! \n", err)
		return
	}

	data.Version = release + "-" + commit

	// Start CLI
	err = cli.Execute(data)
	if err != nil {
		fmt.Println("OMG", err)
	}
}
