package input

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/input/flags"
)

func readFiles(d *data.D) (*data.Config, *data.Config, error) {
	var err error

	path := flags.GetFlagOrEnv(d, "path")
	configFile := path + "/.omg.toml"
	privateFile := path + "/.omg_private.toml"

	// create '.omg.toml' if it does not exist
	err = ensureItExists(configFile)
	if err != nil {
		return nil, nil, err
	}

	// read values in '.omg'
	config, err := readTOML(configFile)
	if err != nil {
		return nil, nil, err
	}

	// merge values in '.omg_private' if it exists
	_, err = os.Stat(privateFile)
	priv := &data.Config{}
	if err == nil {
		priv, err = readTOML(privateFile)
		if err != nil {
			return nil, nil, err
		}
	}

	return config, priv, nil
}

func ensureItExists(file string) error {
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

func readTOML(path string) (*data.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := &data.Config{}
	err = toml.Unmarshal(inputBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
