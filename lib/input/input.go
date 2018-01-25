package input

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Data is a wrapper for Viper. No one outside input should need to import Viper.
type Data struct {
	*viper.Viper
}

// Read parses all input from the outside world into a Data struct
func Read() (*Data, error) {
	d := &Data{viper.New()}

	// get flag data before anything else
	parseFlags(d)

	// add ENV support for any OMG_* variable
	d.SetEnvPrefix("omg")
	d.AutomaticEnv()

	// then save the args
	d.Set("args", pflag.Args())

	// actual read & merge of config values
	d, err := doRead(d)
	if err != nil {
		return d, err
	}

	return d, nil
}
