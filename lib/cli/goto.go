package cli

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

func init() {
	addCommand(gotoCmd)
}

var gotoCmd = &Command{
	Name:  "goto",
	Short: "Open an SSH session with given server",
	Long: `
omg goto [server name]

It opens an SSH session on a new terminal window on the server with given name.
The actual terminal command can be configured as a template.
See the '.omg.toml' file for more detail.
`,
	Run: gotoFunc,
}

func gotoFunc(cmd *Command, data *data.D) error {
	if len(data.Args) < 2 {
		return fmt.Errorf("Not enough arguments. Usage:\n%s", cmd.Long)
	}
	name := data.Args[1]

	servers, err := input.ResolveServers(data)
	if err != nil {
		return err
	}
	data.Config.Servers = servers

	target := data.Config.Servers[name]
	if target == nil {
		return fmt.Errorf("Unrecognised server name: '%s'", name)
	}
	if target.IP == "" {
		return fmt.Errorf("Server without IP: '%+v'", target)
	}
	return openTerminal(target, data)
}

func openTerminal(target *data.Server, data *data.D) error {
	cmdline, err := renderTerminalTemplate(target, data)
	if err != nil {
		return err
	}
	_, err = hlp.Run(hlp.PrintToStdout, cmdline)
	return err
}

func renderTerminalTemplate(target *data.Server, data *data.D) (string, error) {
	strTpl := data.Config.Terminal
	tpl, err := template.New("term").Parse(strTpl)
	if err != nil {
		return "", fmt.Errorf("Bad template for terminal: %s", strTpl)
	}

	tplData := struct {
		Title   string
		Command string
	}{
		target.Name,
		fmt.Sprintf("ssh %s@%s", getRemoteUser(target, data), target.IP),
	}

	var res bytes.Buffer
	err = tpl.Execute(&res, tplData)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func getRemoteUser(target *data.Server, data *data.D) string {
	if target.RemoteUser != "" {
		return target.RemoteUser
	}
	return data.Config.RemoteUser
}
