package input

import (
	"flag"
	"os"
)

// setFileDefaults is where you should code default values file configurable options
func setFileDefaults(d *Data) {
	// d.SetDefault("terminal", "terminator -T '{{.Title}}' -e '{{.Command}}'")
	d.SetDefault("terminal", "konsole -e \"{{.Command}}\"")
	d.SetDefault("remoteUser", "$USER")
}

// defineFlags is where you should code supported flags
func defineFlags(fset *flag.FlagSet) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fset.String("path", cwd, "Folder with configuration files.")

	return nil
}
