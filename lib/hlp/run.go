package hlp

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/rubencaro/omg/lib/data"
)

// Silent would make Run produce no output on stdout
const Silent = 0

// PrintToStdout will make Run print to stdout all output on realtime
const PrintToStdout = 1

// RunOpts is the options struct to pass to Run
type RunOpts struct {
	// Silent or PrintToStdout, to control how output should be produced
	// Default to 0=Silent
	Print int

	// To mark each line of output
	// Default to empty string
	Prefix string
}

// Run gets a bash command string and runs it on a new bash instance.
// It captures its output. If 'print' is 'PrintToStdout' it will also print
// any output to stdout on realtime.
func Run(cmdline string, opts *RunOpts, args ...string) (string, error) {
	var outBuf, errBuf bytes.Buffer
	if opts == nil {
		opts = &RunOpts{}
	}

	args = append([]string{cmdline}, args...)
	cmd := exec.Command("bash", "-c", strings.Join(args, " "))

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	outScanner := bufio.NewScanner(outReader)
	go capture(outScanner, &outBuf, *opts)

	errScanner := bufio.NewScanner(errReader)
	go capture(errScanner, &errBuf, *opts)

	err = cmd.Start()
	if err != nil {
		return outBuf.String() + errBuf.String(), err
	}

	err = cmd.Wait()
	if err != nil {
		return outBuf.String() + errBuf.String(), err
	}

	return outBuf.String() + errBuf.String(), nil
}

func capture(scanner *bufio.Scanner, buf *bytes.Buffer, opts RunOpts) {
	var txt string
	var err error
	for scanner.Scan() {
		txt = scanner.Text()
		if opts.Print == PrintToStdout {
			_, err = fmt.Println(opts.Prefix + txt)
			if err != nil {
				panic(err.Error()) // we cannot go to stdout!
			}
		}
		_, err = buf.WriteString(txt)
		if err != nil {
			fmt.Println("OMG we could not write to byte buffer! \n", err)
			return
		}
	}
}

// RunForEachServer runs given cmdline on each server on given D
// Returns a single error with a detailed message of who failed and why
func RunForEachServer(cmdline string, d *data.D) error {
	var wg sync.WaitGroup
	// an errors channel with enough buffer to hold all responses
	// we want to Wait before reading them
	errors := make(chan error, len(d.Config.Servers))
	for _, s := range d.Config.Servers {
		wg.Add(1)
		go runSingleCmd(&wg, errors, s, cmdline, d)
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

func runSingleCmd(wg *sync.WaitGroup, errors chan<- error, s *data.Server, cmdline string, d *data.D) {
	defer wg.Done()
	if s == nil {
		return
	}
	exports := getSingleExportsString(s, d)
	prefix := fmt.Sprintf("%s(%s): ", s.Name, s.IP)
	_, err := Run(exports+cmdline, &RunOpts{Print: PrintToStdout, Prefix: prefix}, d.Args[1:]...)
	errors <- err
}

// RunRegularCmd run given cmdline once and returns its error
func RunRegularCmd(cmdline string, d *data.D) error {
	exports := getRegularExportsString(d)
	_, err := Run(exports+cmdline, &RunOpts{Print: PrintToStdout}, d.Args[1:]...)
	return err
}

func getRegularExportsString(d *data.D) string {
	names := GetServerNames(d)
	res := "export OMG_SERVER_NAMES=" + strings.Join(names, ",") + ";"

	ips := getServerIPs(d)
	res += "export OMG_SERVER_IPS=" + strings.Join(ips, ",") + ";"

	return res + " "
}

func getSingleExportsString(s *data.Server, d *data.D) string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf(
		"export OMG_SERVER_NAME=%s;export OMG_SERVER_IP=%s;export OMG_USER=%s;",
		s.Name,
		s.IP,
		GetRemoteUser(s, d),
	)
}

// GetServerNames returns server names contained on given D
func GetServerNames(d *data.D) []string {
	res := []string{}
	for _, s := range d.Config.Servers {
		if s == nil {
			continue
		}
		res = append(res, s.Name)
	}
	return res
}

func getServerIPs(d *data.D) []string {
	res := []string{}
	for _, s := range d.Config.Servers {
		if s == nil {
			continue
		}
		res = append(res, s.IP)
	}
	return res
}

// GetRemoteUser returns the remote user for this target
// after looking on every configuration level
func GetRemoteUser(target *data.Server, d *data.D) string {
	if target != nil && target.RemoteUser != "" {
		return target.RemoteUser
	}
	return d.Config.RemoteUser
}
