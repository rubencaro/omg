// Package data contains all common data exchange structs
// needed by many other packages
package data

import "flag"

// D is the main input data struct
type D struct {
	// Version is to be set from build script, saved here to be used around
	Version string
	// Args is the list of non-flag Args
	Args []string
	// FlagSet is the struct for the parsed flags
	FlagSet *flag.FlagSet
	// Config is the struct for the data from .omg.toml
	Config *Config
	// Private is the struct for the data from .omg_private.toml
	Private *Config
	// Defaults are coded defaul values
	Defaults *Config
}
