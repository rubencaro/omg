package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

func init() {
	addCommand(runCmd)
}

var runCmd = &Command{
	Name:  "run",
	Short: "Run given cmdline through SSH on requested servers",
	Long:  getRunUsage(),
	Run:   runFunc,
}

func runFunc(cmd *Command, d *data.D) error {
	servers, err := input.ResolveServers(d)
	if err != nil {
		return err
	}
	d.Config.Servers = servers

	cmdline := "ssh $OMG_USER@$OMG_SERVER_IP "
	ok := hlp.Confirm("This will run '%s'\non %s. \nAre you sure?", cmdline, hlp.GetServerNames(d))
	if !ok {
		return fmt.Errorf("Cancelled")
	}
	return hlp.RunForEachServer(cmdline, d)
}

func getRunUsage() string {
	return fmt.Sprintf(`
omg [target selection flags] run [cmdline]

It runs given cmdline on each selected target server via SSH.
It captures both stdout and stderr and prints them to stdout as they come.

%s
`, getEnvVariablesEachUsage())
}
