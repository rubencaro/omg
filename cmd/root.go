package cmd

import (
	"fmt"
	"os"

	"github.com/rubencaro/omg/lib/cnf"
	"github.com/spf13/cobra"
)

// c is the (private) global containing all config stuff
var c *cnf.Cnf

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "omg",
	Short: "All those helper scripts around your code should be amazing too",
	Long: `A tool to efficiently manage all those little things/scripts/files around my code that make up for different stages of development. Such as compiling, releasing, deploying, packaging or publishing, but also linting, formatting, testing, benchmarking, etc.
	It is stack independent, so I can use it with all my Javascript/Elixir/Go/Whatever projects. Also compatible with all the shapes and colors of helper scripts and tools I use with them.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	conf, err := cnf.Read()
	if err != nil {
		fmt.Println("We need configuration!", err)
		os.Exit(1)
	}
	c = &conf
}
