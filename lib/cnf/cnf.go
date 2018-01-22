//Package cnf encapsulates config processing functions
package cnf

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Cnf is a wrapper for Viper. No one outside cnf should need to import Viper.
type Cnf struct {
	*viper.Viper
}

// Read reads config for this application and returns it.
//
// First it will retrieve 'path' flag. It defaults to the current path.
// Then it reads the mandatory '.omg.toml' file, and then it tries to read
// an optional '.omg_private.toml' file.
//
// You should keep '.omg_private.toml' outside your version control.
//
// A typical use case:
//		var c, err = cnf.Read()
//		if err != nil {
//			fmt.Println("We need configuration!", err)
// 			return
//		}
//
func Read() (*Cnf, error) {
	return ReadAndValidate([]string{}, make(map[string]interface{}))
}

// ReadAndValidate is like Read but adding validation features.
//
// It accepts a 'mandatory' array which should contain the keys that are needed.
// All of those keys must have a value, else it will return an error.
// It also accepts a 'defaults' map, which should contain default values for optional keys.
//
// A typical use case:
// 		var c, err = cnf.ReadAndValidate(
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
func ReadAndValidate(mandatory []string, defaults map[string]interface{}) (*Cnf, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Get minimum config from flags
	var path = flag.String("path", cwd, "Folder with configuration files.")
	flag.Parse()

	return doReadAndValidate(*path, mandatory, defaults)
}

func doReadAndValidate(path string, mandatory []string, defaults map[string]interface{}) (*Cnf, error) {
	// actual read & merge of config values
	cnf, err := doRead(path)
	if err != nil {
		return &Cnf{cnf}, err
	}

	// apply defaults
	for k, v := range defaults {
		cnf.SetDefault(k, v)
	}

	// check mandatory values are set
	for i := 0; i < len(mandatory); i++ {
		if !cnf.IsSet(mandatory[i]) {
			return &Cnf{cnf}, fmt.Errorf("mandatory key '%s' not found in config", mandatory[i])
		}
	}

	return &Cnf{cnf}, nil
}

func doRead(path string) (*viper.Viper, error) {
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

	// create '.omg.toml' if it does not exist
	err = ensureItExists(path)
	if err != nil {
		return v, err
	}

	// read values in '.omg'
	v.SetConfigName(".omg")
	err = v.ReadInConfig()
	if err != nil {
		return v, err
	}

	// merge values in '.omg_private' if exists
	v.SetConfigName(".omg_private")
	err = v.MergeInConfig()
	_, notFound = err.(viper.ConfigFileNotFoundError)
	if err != nil && !notFound {
		return v, err
	}

	return v, nil
}

func ensureItExists(path string) error {
	file := path + "/.omg.toml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return createSample(file)
	}
	return nil
}

func createSample(dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, strings.NewReader(sample))
	return err
}
