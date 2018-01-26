package input

import (
	"os"

	"github.com/spf13/pflag"
)

func parseCmdline(d *Data) error {
	// first parse and bind flags
	err := parseFlags()
	if err != nil {
		return err
	}
	d.BindPFlags(pflag.CommandLine)

	// then get args after flags
	d.Args = pflag.Args()

	return nil
}

func parseFlags() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// define any flag supported by any command
	pflag.String("path", cwd, "Folder with configuration files.")

	pflag.Parse()

	return nil
}
