package cli

import "github.com/rubencaro/omg/lib/input"

// Command is the struct for a CLI command
type Command struct {
	Name string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// The actual running function
	Run func(cmd *Command, data *input.Data) error
}

// Server is the data for a remote server
type Server struct {
	Name string
	IP   string
}
