package cmd

import (
	"fmt"

	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
	"github.com/spf13/cobra"
)

// These are (private) globals
// These are global inside cmd just to go over Cobra
// Commands should pass it as argument to any code outside cmd package
var d *input.Data
var version string // to be set from build script

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "omg",
	Short: "OMG is a tool to manage all those little scripts around code",
	Long: `OMG is a tool to efficiently manage all those little things/scripts/files
around my code that make up for different stages of development.
Such as compiling, releasing, deploying, packaging or publishing,
but also linting, formatting, testing, benchmarking, etc.

See the generated '.omg.toml' file for configuration options.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(data *input.Data, ver string) error {
	// save things into private globals
	d = data
	version = ver

	// Add whatever commands may be inside config
	err := addDyamicCommands(d, rootCmd)
	if err != nil {
		return err
	}

	return rootCmd.Execute()
}

func addDyamicCommands(d *input.Data, root *cobra.Command) error {
	for k, v := range d.GetStringMapString("custom") {
		addCommand(k, v, root)
	}
	return nil
}

func addCommand(name, cmdline string, root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Just run '%s'", cmdline),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := hlp.Run(cmdline, args...)
			return err
		},
		DisableFlagParsing: true,
	}
	root.AddCommand(cmd)
}
