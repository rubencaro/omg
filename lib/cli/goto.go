package cli

import (
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

func init() {
	addCommand(gotoCmd)
}

var gotoCmd = &Command{
	Name:  "goto",
	Short: "Open an SSH session with given server",
	Long: `It opens an SSH session on a new terminal window on the server with given name.
	The actual "terminal" command can be configured as a template.`,
	Run: func(cmd *Command, data *input.Data) error {
		hlp.Spit("hey")
		return nil
	},
}
