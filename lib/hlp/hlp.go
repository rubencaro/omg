// Package hlp contains useful universal helpers
// Keep it small or it will be a smell
package hlp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// Spit prints anything given to stdout
func Spit(what interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\n\033[1;91m%s:%d\n%+v\n\n\033[00m", file, line, what)
}

// Confirm asks the user for confirmation. A user must type in "yes" or anything else and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it returns false.
func Confirm(msg string, args ...interface{}) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(msg+" [y/N]: ", args...)

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.ToLower(strings.TrimSpace(response))

	if response == "y" || response == "yes" {
		return true
	}
	return false
}
