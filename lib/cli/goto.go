package cli

import (
	"bytes"
	"fmt"
	"text/template"

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

func gotoFunc(cmd *Command, data *input.Data) error {
	if len(data.Args) < 2 {
		return fmt.Errorf("Not enough arguments. Usage:\n%s", cmd.Long)
	}

	raw := data.GetStringMap("servers." + data.Args[1])
	var target = &Server{Name: data.Args[1], IP: raw["ip"].(string)}
	return openTerminal(target, data)
}

func openTerminal(target *Server, data *input.Data) error {
	cmdline, err := renderTerminalTemplate(target, data)
	if err != nil {
		return err
	}
	hlp.Spit(cmdline)

	_, err = hlp.Run(cmdline)
	return err
}

func renderTerminalTemplate(target *Server, data *input.Data) (string, error) {
	strTpl := data.GetString("terminal")
	tpl, err := template.New("term").Parse(strTpl)
	if err != nil {
		return "", fmt.Errorf("Bad template for terminal: %s", strTpl)
	}

	tplData := struct {
		Title   string
		Command string
	}{
		target.Name,
		fmt.Sprintf("ssh %s@%s", data.GetString("remoteUser"), target.IP),
	}

	var res bytes.Buffer
	err = tpl.Execute(&res, tplData)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
