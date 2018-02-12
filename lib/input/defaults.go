package input

import (
	"flag"
	"os"

	"github.com/rubencaro/omg/lib/data"
)

// getDefaults is where you should code default values for non-flag options
func getDefaults() *data.Config {
	return &data.Config{
		Terminal:   "konsole -e '{{.Command}}' &",
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
	fset.String("servers", "", "Comma separated list of server names (ex. 'srv1,srv2,srv3').\n        Overrides other ways to determine the target server list.")
	fset.String("match", "", "Regular expression to be matched agaist the list of server names.\n        Applies regardless of the way to obtain that list.")

	return nil
}
