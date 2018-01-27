package input

import (
	"flag"
	"os"

	"github.com/spf13/viper"
)

// Data is a wrapper for Viper. No one outside input should need to import Viper.
type Data struct {
	*viper.Viper
	Version string // to be set from build script, saved here to be used around
	Args    []string
	FlagSet *flag.FlagSet
}

// Get returns the value for given key.
// It looks up for it on flags first, then env, and then config files.
// Every key has at least a default value when it's defined.
// Get expects the requested key to have at least a default value.
// Either meaning it should be defined on code somewhere but it's not,
// or requested key was incorrect, Get will panic when it cannot find a value.
func Get(d *Data, key string) string {
	// try flags first
	v, ok := getFlag(d, key)
	if ok {
		return v
	}

	// env second
	v, ok = os.LookupEnv("OMG_" + key)
	if ok {
		return v
	}

	// files last
	// TODO: remove this with proper parsed toml struct access
	return d.GetString(key)

	// panic(fmt.Sprintf("OMG Unexpected key '%s'", key))
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
	d := &Data{Viper: viper.New()}

	// then get cmdline data
	parseCmdline(d)

	// actual read & merge of config values
	setFileDefaults(d)
	d, err := doRead(d)
	if err != nil {
		return d, err
	}

	return d, nil
}
