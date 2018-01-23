// Package hlp contains useful universal helpers
// Keep it small or it will be a smell
package hlp

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Spit prints anything given to stdout
func Spit(what interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\n\033[1;91m%s:%d\n%+v\n\n\033[00m", file, line, what)
}

// Run gets a bash command string and runs it on a new bash instance.
// It captures its output and prints it to stdout.
//
func Run(cmdline string, args ...string) error {
	args = append([]string{cmdline}, args...)
	cmd := exec.Command("bash", "-c", strings.Join(args, " "))

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()

	errReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
