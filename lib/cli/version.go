package cli

import (
	"fmt"

	"github.com/rubencaro/omg/lib/input"
)

func init() {
	addCommand(versionCmd)
}

var versionCmd = &Command{
	Name:  "version",
	Short: "Print the version number of OMG",
	Long: `It's composed of a numerical part X.Y.Z (Major.Minor.Patch)
followed by a pre-release version.

Just take a look at https://semver.org/.`,
	Run: func(cmd *Command, data *input.Data) error {
		fmt.Printf("0.1.0-%s\n", data.Version)
		return nil
	},
}
