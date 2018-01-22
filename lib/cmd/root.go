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
	Use: "omg",
	// Short: "OMG is a tool to manage all those little scripts around code",
	// Long: `
	// A tool to efficiently manage all those little things/scripts/files
	// around my code that make up for different stages of development.
	// Such as compiling, releasing, deploying, packaging or publishing,
	// but also linting, formatting, testing, benchmarking, etc.

	// Because all those helper scripts around your code should be amazing too.

	// It is stack independent, so I can use it with all my
	// Javascript/Elixir/Go/Whatever projects. Also compatible with all the
	// shapes and colors of helper scripts and tools I use with them.
	// `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(conf *cnf.Cnf) error {
	// save config into the private global
	c = conf

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
