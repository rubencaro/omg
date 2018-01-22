package cmd

import (
	"fmt"

	"github.com/rubencaro/omg/lib/cnf"
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/spf13/cobra"
)

// c is the (private) global containing all config stuff
// commands should pass it as argument to any code outside this package
var c *cnf.Cnf

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "omg",
	Short: "OMG is a tool to manage all those little scripts around code",
	Long: `
 A tool to efficiently manage all those little things/scripts/files
 around my code that make up for different stages of development.
 Such as compiling, releasing, deploying, packaging or publishing,
 but also linting, formatting, testing, benchmarking, etc.

 Because all those helper scripts around your code should be amazing too.

 It is stack independent, so I can use it with all my
 Javascript/Elixir/Go/Whatever projects. Also compatible with all the
 shapes and colors of helper scripts and tools I use with them.
`,
	Run: func(cmd *cobra.Command, args []string) {
		customCmd := c.GetString("custom." + args[0])
		if customCmd != "" {
			hlp.Run(c.GetString("custom." + args[0]))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	// save config to the global c to make it accessible to everyone in cmd
	err := initConfig()
	if err != nil {
		return err
	}

	return rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	conf, err := cnf.Read()
	if err != nil {
		return fmt.Errorf("We need configuration! %v", err)
	}
	c = conf
	return nil
}
