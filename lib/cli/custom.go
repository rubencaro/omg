package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

func addCustomCommands(d *data.D) error {
	for k, v := range d.Config.Customs {
		addCustomCommand(k, v, d)
	}
	return nil
}

func addCustomCommand(name string, cust *data.Custom, d *data.D) {
	c := &Command{
		Name:  name,
		Short: fmt.Sprintf("Just run '%s'", cust.Cmd),
		Long:  getCustomUsage(name, cust),
		Run:   customFunc(cust),
	}
	addCommand(c)
}

func customFunc(cust *data.Custom) func(*Command, *data.D) error {
	return func(cmd *Command, d *data.D) error {
		servers, err := input.ResolveServers(d)
		if err != nil {
			return err
		}
		d.Config.Servers = servers

		if cust.Run == "each" {
			ok := hlp.Confirm("This will run '%s'\non %s. \nAre you sure?", cust.Cmd, hlp.GetServerNames(d))
			if !ok {
				return fmt.Errorf("Cancelled")
			}
			return hlp.RunForEachServer(cust.Cmd, d)
		}
		return hlp.RunRegularCmd(cust.Cmd, d)
	}
}

func getCustomUsage(name string, cust *data.Custom) string {
	return fmt.Sprintf(`
omg %s [...]

It will run '%s [...]'.

Just as configured in the 'custom' section of the '.omg.toml' file.
Following arguments and flags will be passed along as given.

The following environment variables will be set for you:

%s
		`, name, cust.Cmd, getEnvVariablesUsage(cust))
}

func getEnvVariablesUsage(cust *data.Custom) string {
	if cust.Run == "each" {
		return getEnvVariablesEachUsage()
	}
	return `
OMG_SERVER_NAMES - will contain selected servers' names
                   comma separated (ex. "srv1,srv2")
OMG_SERVER_IPS   - will contain selected servers' IPs
									 comma separated (ex. "1.1.1.1,1.1.1.2")

The comma separated format is suited to pass on as part of an URL,
and dead easy to split into an array on most programming languages.
For example if you need to use it from bash you can get an array
doing something like:

		arrServerNames=(${OMG_SERVER_NAMES//,/ })
		`
}

func getEnvVariablesEachUsage() string {
	return `
OMG_SERVER_NAME - will contain selected server's name (ex. "srv1")
OMG_SERVER_IP   - will contain selected server's IP (ex. "1.1.1.1")
OMG_USER        - will contain the configured remote user
`
}
