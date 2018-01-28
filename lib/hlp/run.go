package hlp

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Silent would make Run produce no output on stdout
const Silent = 0

// PrintToStdout will make Run print to stdout all output on realtime
const PrintToStdout = 1

// Run gets a bash command string and runs it on a new bash instance.
// It captures its output. If 'print' is 'PrintToStdout' it will also print
// any output to stdout on realtime.
func Run(print int, cmdline string, args ...string) (string, error) {
	var outBuf, errBuf bytes.Buffer

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
	go capture(outScanner, &outBuf, print)

	errScanner := bufio.NewScanner(errReader)
	go capture(errScanner, &errBuf, print)

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

func capture(scanner *bufio.Scanner, buf *bytes.Buffer, print int) {
	var txt string
	var err error
	for scanner.Scan() {
		txt = scanner.Text()
		if print == PrintToStdout {
			_, err = fmt.Println(txt)
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
