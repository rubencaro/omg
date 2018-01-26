package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/input"
)

// holder for init-time definition of commands
var commands = map[string]*Command{}

func addCommand(cmd *Command) {
	commands[cmd.Name] = cmd
}

// Execute finds out which Command to run, and then runs it
func Execute(data *input.Data) error {
	if len(data.Args) == 0 {
		return fmt.Errorf("At least command name is needed")
	}

	cmd := commands[data.Args[0]]
	return cmd.Run(cmd, data)
}
