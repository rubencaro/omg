package input

import (
	"flag"
	"os"
)

func parseCmdline(d *Data) error {
	// first parse and bind flags
	fset, err := parseFlags()
	if err != nil {
		return err
	}
	d.FlagSet = fset

	// then get args after flags
	d.Args = fset.Args()

	return nil
}

func parseFlags() (*flag.FlagSet, error) {
	fset := flag.NewFlagSet("main", flag.ContinueOnError)

	err := defineFlags(fset)
	if err != nil {
		return nil, err
	}

	fset.Parse(os.Args)

	return fset, nil
}
