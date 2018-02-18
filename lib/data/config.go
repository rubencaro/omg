package data

// Here are defined all the structs to validate the config file.

// Config is the base struct for the config file
type Config struct {
	// Terminal is the teplate for the actual command to be run to open a terminal
	Terminal string

	// RemoteUser is the user to be used on remote machines
	RemoteUser string

	// Customs is the definition of custom commands to be mapped through OMG
	Customs map[string]*Custom

	// Servers is the manual server list to work with
	Servers map[string]*Server

	// Gce is the config to get the server list from GCE
	// If this is set, the 'Servers' list will be overwritten
	Gce *Gce
}

// Custom is the data defining a custom command
type Custom struct {
	Cmd string
	Run string
}

// Server is the data for a remote server
type Server struct {
	Name       string
	IP         string
	RemoteUser string
}

// Gce is the configuration for GCE services
type Gce struct {
	Project string
	Match   string
}
