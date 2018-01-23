package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OMG",
	Long: `It's composed of a numerical part X.Y.Z (Major.Minor.Patch)
followed by a pre-release version.

Just take a look at https://semver.org/.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("0.1.0-%s\n", version)
	},
}
