package hlp

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Run gets a bash command string and runs it on a new bash instance.
// It captures its output and prints it to stdout.
//
func Run(cmdline string, args ...string) (string, error) {
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
	go capture(outScanner, &outBuf)

	errScanner := bufio.NewScanner(errReader)
	go capture(errScanner, &errBuf)

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

func capture(scanner *bufio.Scanner, buf *bytes.Buffer) {
	var txt string
	var err error
	for scanner.Scan() {
		txt = scanner.Text()
		_, err = fmt.Println(txt)
		if err != nil {
			panic(err.Error()) // we cannot go to stdout!
		}
		_, err = buf.WriteString(txt)
		if err != nil {
			fmt.Println("OMG we could not write to byte buffer! \n", err)
			return
		}
	}
}
