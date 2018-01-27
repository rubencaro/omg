package input

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func doRead(d *Data) (*Data, error) {
	var err error
	var notFound bool

	path := Get(d, "path") // set via flag or env
	d.AddConfigPath(path)

	// create '.omg.toml' if it does not exist
	err = ensureItExists(path)
	if err != nil {
		return d, err
	}

	// read values in '.omg'
	d.SetConfigName(".omg")
	err = d.ReadInConfig()
	if err != nil {
		return d, err
	}

	// merge values in '.omg_private' if exists
	d.SetConfigName(".omg_private")
	err = d.MergeInConfig()
	_, notFound = err.(viper.ConfigFileNotFoundError)
	if err != nil && !notFound {
		return d, err
	}

	return d, nil
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
