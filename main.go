package main

import (
	"fmt"

	"github.com/rubencaro/omg/lib/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("OMG it failed! => ", err)
	}
}
