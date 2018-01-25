package input

import (
	"os"

	"github.com/spf13/pflag"
)

// defineFlags defines all needed flags to be parsed from commandline
func parseFlags(d *Data) error {
	// Get minimum config from flags
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// here define any flag supported by any command
	pflag.String("path", cwd, "Folder with configuration files.")

	pflag.Parse()
	d.BindPFlags(pflag.CommandLine)

	return nil
}
