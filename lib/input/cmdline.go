package input

import (
	"flag"
	"os"
)

func getFlagsAndArgs() (*flag.FlagSet, []string, error) {
	// first parse and bind flags
	fset, err := parseFlags()
	if err != nil {
		return nil, nil, err
	}
	return fset, fset.Args(), nil
}

func parseFlags() (*flag.FlagSet, error) {
	fset := flag.NewFlagSet("OMG flags", flag.ContinueOnError)

	err := defineFlags(fset)
	if err != nil {
		return nil, err
	}

	fset.Parse(os.Args[1:])

	return fset, nil
}
