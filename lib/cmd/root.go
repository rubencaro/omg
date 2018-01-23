package cmd

import (
	"github.com/rubencaro/omg/lib/cnf"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/spf13/cobra"
)

// These are (private) globals
// These are global inside cmd just to go over Cobra
// Commands should pass it as argument to any code outside cmd package
var c *cnf.Cnf
var version string // to be set from build script

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "omg",
	Short: "OMG is a tool to manage all those little scripts around code",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(conf *cnf.Cnf, ver string) error {
	// save things into private globals
	c = conf
	version = ver

	// Add whatever commands may be inside config
	err := addDyamicCommands(c, rootCmd)
	if err != nil {
		return err
	}

	return rootCmd.Execute()
}

func addDyamicCommands(c *cnf.Cnf, root *cobra.Command) error {
	for k, v := range c.GetStringMapString("custom") {
		addCommand(k, v, root)
	}
	return nil
}

func addCommand(name, cmdline string, root *cobra.Command) {
	cmd := &cobra.Command{
		Use: name,
		RunE: func(cmd *cobra.Command, args []string) error {
			return hlp.Run(cmdline)
		},
	}
	root.AddCommand(cmd)
}
