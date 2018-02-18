package cli

import (
	"fmt"
	"strings"
	"sync"

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
		Long:  getUsageString(name, cust),
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
			return runForEachServer(cust, d)
		}
		return runRegularCmd(cust, d)
	}
}

func runForEachServer(cust *data.Custom, d *data.D) error {
	var wg sync.WaitGroup
	// an errors channel with enough buffer to hold all responses
	// we want to Wait before reading them
	errors := make(chan error, len(d.Config.Servers))
	for _, s := range d.Config.Servers {
		wg.Add(1)
		go runSingleCmd(&wg, errors, s, cust, d)
	}
	wg.Wait()
	return evalErrors(errors)
}

func evalErrors(in <-chan error) error {
	errstrs := readErrors(in)
	if len(errstrs) == 0 {
		return nil
	}
	msg := "These commands did fail:\n\n" + strings.Join(errstrs, "\n")
	return fmt.Errorf(msg)
}

func readErrors(in <-chan error) []string {
	var e error
	errors := []string{}
	// after Wait, noone is writing
	for len(in) > 0 {
		e = <-in
		if e != nil {
			errors = append(errors, e.Error())
		}
	}
	return errors
}

func runSingleCmd(wg *sync.WaitGroup, errors chan<- error, s *data.Server, cust *data.Custom, d *data.D) {
	defer wg.Done()
	exports := getSingleExportsString(s)
	_, err := hlp.Run(hlp.PrintToStdout, exports+cust.Cmd, d.Args[1:]...)
	errors <- err
}

func getUsageString(name string, cust *data.Custom) string {
	return fmt.Sprintf(`
omg %s [...]

It will run '%s [...]'.

Just as configured in the 'custom' section of the '.omg.toml' file.
Following arguments and flags will be passed along as given.

The following environment variables will be set for you:

%s
		`, name, cust.Cmd, getEnvVariablesUsageString(cust))
}

func getEnvVariablesUsageString(cust *data.Custom) string {
	if cust.Run == "each" {
		return `
OMG_SERVER_NAME - will contain selected server's name (ex. "srv1")
OMG_SERVER_IP   - will contain selected server's IP (ex. "1.1.1.1")
			`
	}
	return `
OMG_SERVER_NAMES - will contain selected servers' names comma separated (ex. "srv1,srv2")
OMG_SERVER_IPS   - will contain selected servers' IPs comma separated (ex. "1.1.1.1,1.1.1.2")
		`
}

func runRegularCmd(cust *data.Custom, d *data.D) error {
	exports := getRegularExportsString(d)
	_, err := hlp.Run(hlp.PrintToStdout, exports+cust.Cmd, d.Args[1:]...)
	return err
}

func getRegularExportsString(d *data.D) string {
	names := getServerNames(d)
	res := "export OMG_SERVER_NAMES=" + strings.Join(names, ",") + ";"

	ips := getServerIPs(d)
	res += "export OMG_SERVER_IPS=" + strings.Join(ips, ",") + ";"

	return res + " "
}

func getSingleExportsString(s *data.Server) string {
	return fmt.Sprintf("export OMG_SERVER_NAME=%s;export OMG_SERVER_IP=%s;", s.Name, s.IP)
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
