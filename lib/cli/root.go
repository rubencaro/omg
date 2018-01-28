package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
)

// Command is the struct for a CLI command
type Command struct {
	Name string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// The actual running function
	Run func(cmd *Command, data *data.D) error
}

// holder for init-time definition of commands
var commands = map[string]*Command{}

func addCommand(cmd *Command) {
	commands[cmd.Name] = cmd
}

// Execute finds out which Command to run, and then runs it
func Execute(data *data.D) error {
	addDynamicCommands(data)

	var cmd *Command
	if len(data.Args) == 0 { // no command given
		cmd = helpCmd
	} else {
		cmd = commands[data.Args[0]]
	}

	if cmd == nil {
		return fmt.Errorf("Unknown command '%s'", data.Args[0])
	}

	return cmd.Run(cmd, data)
}

func addDynamicCommands(d *data.D) error {
	for k, v := range d.Config.Custom {
		addDynamicCommand(k, v)
	}
	return nil
}

func addDynamicCommand(name, cmdline string) {
	c := &Command{
		Name:  name,
		Short: fmt.Sprintf("Just run '%s'", cmdline),
		Long: fmt.Sprintf(`
omg %s [...]

It will run '%s [...]'.

Just as configured in the 'custom' section of the '.omg.toml' file.
Following arguments and flags will be passed along as given.
		`, name, cmdline),
		Run: func(cmd *Command, data *data.D) error {
			_, err := hlp.Run(hlp.PrintToStdout, cmdline, data.Args[1:]...)
			return err
		},
	}
	addCommand(c)
}
