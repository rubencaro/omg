package input

import (
	"flag"
	"os"
)

// getDefaults is where you should code default values for non-flag options
func getDefaults() *ConfigData {
	return &ConfigData{
		// d.SetDefault("terminal", "terminator -T '{{.Title}}' -e '{{.Command}}'")
		Terminal:   "konsole -e \"{{.Command}}\"",
		RemoteUser: "$USER",
	}
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
