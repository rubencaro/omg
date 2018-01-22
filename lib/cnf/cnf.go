//Package cnf encapsulates config processing functions
package cnf

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// Cnf is a wrapper for Viper. No one outside cnf should need to import Viper.
type Cnf struct {
	*viper.Viper
}

// Read reads config for this application and returns it.
//
// First it will retrieve 'env' and 'config' flags. They will default to 'dev' and './config'.
// Then it uses those path and environment to load files found there.
// It supports anything Viper does (will get TOML, YAML, JSON...).
// It first reads 'config' file, then it reads given 'env' file.
// Then it tries to read a 'private' file.
// Only 'config' file must exist. Others are optional.
//
// A typical config files layout is similar to Elixir's:
//     config/config.toml
//           /prod.toml
//           /dev.toml
//           /test.toml
//           /private.toml  <-- typically on gitignore
//
//
// A typical use case:
//		var cnf, err = cnf.Read()
//		if err != nil {
//			fmt.Println("We need configuration!", err)
// 			return
//		}
//
func Read() (Cnf, error) {
	return ReadAndValidate([]string{}, make(map[string]interface{}))
}

// ReadAndValidate is like Read but adding validation features.
//
// It accepts a 'mandatory' array which should contain the keys that are needed.
// All of those keys must have a value, else it will return an error.
// It also accepts a 'defaults' map, which should contain default values for optional keys.
//
// A typical use case:
// 		var cnf, err = cnf.ReadAndValidate(
// 			[]string{
// 				"mandatorykey1",
// 				"mandatorykey2",
// 			}, map[string]interface{}{
// 				"optionalkey1": "defaultvalue1",
// 			}
//		)
// 		if err != nil {
// 			fmt.Println("We need configuration!", err)
// 			return
// 		}
//
func ReadAndValidate(mandatory []string, defaults map[string]interface{}) (Cnf, error) {
	// Get minimum config from flags
	var env = flag.String("env", "dev", "Environment. Should be something like 'dev', 'prod' or 'test'.")
	var path = flag.String("config", "./config", "Folder with configuration files.")
	flag.Parse()

	return doReadAndValidate(*path, *env, mandatory, defaults)
}

func doReadAndValidate(path string, env string, mandatory []string, defaults map[string]interface{}) (Cnf, error) {
	// actual read & merge of config values
	cnf, err := doRead(path, env)
	if err != nil {
		return Cnf{cnf}, err
	}

	// apply defaults
	for k, v := range defaults {
		cnf.SetDefault(k, v)
	}

	// check mandatory values are set
	for i := 0; i < len(mandatory); i++ {
		if !cnf.IsSet(mandatory[i]) {
			return Cnf{cnf}, fmt.Errorf("mandatory key '%s' not found in config", mandatory[i])
		}
	}

	return Cnf{cnf}, nil
}

func doRead(path string, env string) (*viper.Viper, error) {
	var err error
	var notFound bool

	v := viper.New()

	// add ENV support for any OMG_* variable
	v.SetEnvPrefix("omg")
	v.AutomaticEnv()

	// force config path if given via OMG_PATH
	if v.IsSet("path") {
		path = v.GetString("path")
	}
	v.AddConfigPath(path)

	// read values in 'config'
	v.SetConfigName("config")
	err = v.ReadInConfig()
	if err != nil {
		return v, err
	}

	// force env if given via OMG_ENV
	if v.IsSet("env") {
		env = v.GetString("env")
	}
	// merge values in 'env' if exists
	v.SetConfigName(env)
	err = v.MergeInConfig()
	_, notFound = err.(viper.ConfigFileNotFoundError)
	if err != nil && !notFound {
		return v, err
	}

	// merge values in 'private' if exists
	v.SetConfigName("private")
	err = v.MergeInConfig()
	_, notFound = err.(viper.ConfigFileNotFoundError)
	if err != nil && !notFound {
		return v, err
	}

	return v, nil
}
