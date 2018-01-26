package input

import (
	"github.com/spf13/viper"
)

// Data is a wrapper for Viper. No one outside input should need to import Viper.
type Data struct {
	*viper.Viper
	Version string // to be set from build script, saved here to be used around
	Args    []string
}

// Read parses all input from the outside world into a Data struct
func Read() (*Data, error) {
	d := &Data{Viper: viper.New()}

	// get cmdline data before anything else
	parseCmdline(d)

	// add ENV support for any OMG_* variable
	d.SetEnvPrefix("omg")
	d.AutomaticEnv()

	// actual read & merge of config values
	d, err := doRead(d)
	if err != nil {
		return d, err
	}

	return d, nil
}
