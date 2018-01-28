package input

import (
	"flag"
	"os"

	"github.com/rubencaro/omg/lib/data"
)

// getDefaults is where you should code default values for non-flag options
func getDefaults() *data.Config {
	return &data.Config{
		Terminal:   "konsole -e '{{.Command}}' &",
		RemoteUser: "$USER",
	}
}

// defineFlags is where you should code supported flags
func defineFlags(fset *flag.FlagSet) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fset.String("path", cwd, "Folder with configuration files.")

	return nil
}
