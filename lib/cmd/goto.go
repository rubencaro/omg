package cmd

import (
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gotoCmd)
}

var gotoCmd = &cobra.Command{
	Use:   "goto [server]",
	Short: "Open an SSH session with given server",
	Long: `It opens an SSH session on a new terminal window on the server with given name.
	The actual "terminal" command can be configured as a template.`,
	Run: func(cmd *cobra.Command, args []string) {
		hlp.Spit("hey")
	},
}
