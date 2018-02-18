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

func gotoFunc(cmd *Command, d *data.D) error {
	if len(d.Args) < 2 {
		return fmt.Errorf("Not enough arguments. Usage:\n%s", cmd.Long)
	}
	name := d.Args[1]

	servers, err := input.ResolveServers(d)
	if err != nil {
		return err
	}
	d.Config.Servers = servers

	target := d.Config.Servers[name]
	if target == nil {
		return fmt.Errorf("Unrecognised server name: '%s'", name)
	}
	if target.IP == "" {
		return fmt.Errorf("Server without IP: '%+v'", target)
	}
	return openTerminal(target, d)
}

func openTerminal(target *data.Server, d *data.D) error {
	cmdline, err := renderTerminalTemplate(target, d)
	if err != nil {
		return err
	}
	_, err = hlp.Run(cmdline, &hlp.RunOpts{Print: hlp.PrintToStdout})
	return err
}

func renderTerminalTemplate(target *data.Server, d *data.D) (string, error) {
	strTpl := d.Config.Terminal
	tpl, err := template.New("term").Parse(strTpl)
	if err != nil {
		return "", fmt.Errorf("Bad template for terminal: %s", strTpl)
	}

	tplData := struct {
		Title   string
		Command string
	}{
		target.Name,
		fmt.Sprintf("ssh %s@%s", hlp.GetRemoteUser(target, d), target.IP),
	}

	var res bytes.Buffer
	err = tpl.Execute(&res, tplData)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
