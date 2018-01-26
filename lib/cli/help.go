package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/input"
)

func init() {
	addCommand(helpCmd)
}

var helpCmd = &Command{
	Name:  "help",
	Short: "Show help about other commands",
	Long:  helpLongText,
	Run: func(data *input.Data) error {
		if len(data.Args) < 2 { // 'help' itself, or no command, given
			printHelpIndex()
		} else {
			printLongHelp(commands[data.Args[1]])
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

func printLongHelp(cmd *Command) {
	if cmd == nil {
		printHelpIndex()
	} else {
		fmt.Println(cmd.Long)
	}
}
