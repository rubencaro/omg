package input

import (
	"flag"
	"os"
)

// Data is the main input data struct
type Data struct {
	// Version is to be set from build script, saved here to be used around
	Version string
	// Args is the list of non-flag Args
	Args []string
	// FlagSet is the struct for the parsed flags
	FlagSet *flag.FlagSet
	// Config is the struct for the data from .omg.toml
	Config *ConfigData
	// Private is the struct for the data from .omg_private.toml
	Private *ConfigData
	// Defaults are coded defaul values
	Defaults *ConfigData
}

// GetFlagOrEnv looks for given key on flags, and if it doesn't find a value,
// then it looks for a 'OMG_<key>' environment variable.
// Returns empty string if not found.
func GetFlagOrEnv(d *Data, key string) string {
	v, ok := getFlag(d, key)
	if ok {
		return v
	}
	return os.Getenv("OMG_" + key)
}

func getFlag(d *Data, key string) (string, bool) {
	f := d.FlagSet.Lookup(key)
	if f == nil {
		return "", false
	}
	return f.Value.String(), true
}

// Read parses all input from the outside world into a Data struct
func Read() (*Data, error) {
	d := &Data{}
	d.Defaults = getDefaults()

	// get cmdline data
	fset, args, err := getFlagsAndArgs()
	if err != nil {
		return d, err
	}
	d.FlagSet = fset
	d.Args = args

	// actual read of config files
	config, priv, err := readFiles(d)
	if err != nil {
		return d, err
	}
	d.Config = config
	d.Private = priv

	// put everything together on the same struct in order of precedence
	d.Config = consolidateData(d)

	return d, nil
}
