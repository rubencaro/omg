package cli

import (
	"flag"
	"fmt"

	"github.com/rubencaro/omg/lib/data"
)

func init() {
	addCommand(helpCmd)
}

var helpCmd = &Command{
	Name:  "help",
	Short: "Show help about other commands",
	Long:  helpLongText,
	Run: func(cmd *Command, d *data.D) error {
		if len(d.Args) < 2 { // 'help' itself, or no command, given
			printHelpIndex()
		} else {
			printLongHelp(commands[d.Args[1]], d.FlagSet)
		}
		return nil
	},
}

var helpLongText = `
Try 'omg help [command]' for more detail about a specific command.
Take a look into the selfgenerated '.omg.toml' file to know more about
configuration options.
`

func printHelpIndex() {
	fmt.Printf("OMG helps you stay amazed. Available commands are:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s\t%s\n", cmd.Name, cmd.Short)
	}
	fmt.Println(helpLongText)
}

func printLongHelp(cmd *Command, fset *flag.FlagSet) {
	if cmd == nil {
		printHelpIndex()
	} else {
		fmt.Println(cmd.Long)
		fset.Usage()
		fmt.Println("")
	}
}
