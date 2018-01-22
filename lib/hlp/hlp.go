// Package hlp contains useful universal helpers
// Keep it small or it will be a smell
package hlp

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Spit prints anything given to stdout
func Spit(what interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\n\033[1;91m%s:%d\n%+v\n\n\033[00m", file, line, what)
}

// Run gets a bash command string and runs it on a new bash instance.
// Then returns whatever `exec.Run` returns.
//
func Run(cmd string) error {
	return exec.Command("bash", "-c", cmd).Run()
}
