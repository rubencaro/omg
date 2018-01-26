package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

// holder for init-time definition of commands
var commands = map[string]*Command{}

func addCommand(cmd *Command) {
	commands[cmd.Name] = cmd
}

// Execute finds out which Command to run, and then runs it
func Execute(data *input.Data) error {
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

	return cmd.Run(data)
}

func addDynamicCommands(d *input.Data) error {
	for k, v := range d.GetStringMapString("custom") {
		addDynamicCommand(k, v)
	}
	return nil
}

func addDynamicCommand(name, cmdline string) {
	cmd := &Command{
		Name:  name,
		Short: fmt.Sprintf("Just run '%s'", cmdline),
		Long: fmt.Sprintf(`
omg %s [...]

It will run '%s [...]'.

Just as configured in the 'custom' section of the '.omg.toml' file.
Following arguments and flags will be passed along as given.
		`, name, cmdline),
		Run: func(data *input.Data) error {
			_, err := hlp.Run(cmdline, data.Args[1:]...)
			return err
		},
	}
	addCommand(cmd)
}
