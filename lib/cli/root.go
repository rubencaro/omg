package cli

import (
	"fmt"
	"strings"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
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
		addDynamicCommand(k, v, d)
	}
	return nil
}

func addDynamicCommand(name, cmdline string, d *data.D) {

	c := &Command{
		Name:  name,
		Short: fmt.Sprintf("Just run '%s'", cmdline),
		Long: fmt.Sprintf(`
omg %s [...]

It will run '%s [...]'.

Just as configured in the 'custom' section of the '.omg.toml' file.
Following arguments and flags will be passed along as given.

The following environment variables will be set for you:

OMG_SERVER_NAMES - will contain selected servers' names comma separated (ex. "srv1,srv2")
OMG_SERVER_IPS   - will contain selected servers' IPs comma separated (ex. "1.1.1.1,1.1.1.2")

			`, name, cmdline),
		Run: func(cmd *Command, d *data.D) error {
			servers, err := input.ResolveServers(d)
			if err != nil {
				return err
			}
			d.Config.Servers = servers

			exports := getExportsString(d)

			_, err = hlp.Run(hlp.PrintToStdout, exports+cmdline, d.Args[1:]...)
			return err
		},
	}
	addCommand(c)
}

func getExportsString(d *data.D) string {
	names := getServerNames(d)
	res := "export OMG_SERVER_NAMES=" + strings.Join(names, ",") + ";"

	ips := getServerIPs(d)
	res += "export OMG_SERVER_IPS=" + strings.Join(ips, ",") + ";"

	return res + " "
}

func getServerNames(d *data.D) []string {
	res := []string{}
	for _, s := range d.Config.Servers {
		res = append(res, s.Name)
	}
	return res
}

func getServerIPs(d *data.D) []string {
	res := []string{}
	for _, s := range d.Config.Servers {
		res = append(res, s.IP)
	}
	return res
}
