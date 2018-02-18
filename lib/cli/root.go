package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/data"
)

// Command is the struct for a CLI command
type Command struct {
	Name string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// The actual running function
	Run func(*Command, *data.D) error
}

// holder for init-time definition of commands
var commands = map[string]*Command{}

func addCommand(cmd *Command) {
	commands[cmd.Name] = cmd
}

// Execute finds out which Command to run, and then runs it
func Execute(d *data.D) error {
	addCustomCommands(d)

	var cmd *Command
	if len(d.Args) == 0 { // no command given
		cmd = helpCmd
	} else {
		cmd = commands[d.Args[0]]
	}

	if cmd == nil {
		return fmt.Errorf("Unknown command '%s'", d.Args[0])
	}

	return cmd.Run(cmd, d)
}
