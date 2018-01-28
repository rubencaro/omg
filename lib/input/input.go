package input

import (
	"os"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/input/gcloud"
)

// getFlagOrEnv looks for given key on flags, and if it doesn't find a value,
// then it looks for a 'OMG_<key>' environment variable.
// Returns empty string if not found.
func getFlagOrEnv(d *data.D, key string) string {
	v, ok := getFlag(d, key)
	if ok {
		return v
	}
	return os.Getenv("OMG_" + key)
}

func getFlag(d *data.D, key string) (string, bool) {
	f := d.FlagSet.Lookup(key)
	if f == nil {
		return "", false
	}
	return f.Value.String(), true
}

// Read parses all input from the outside world into a data.D struct
func Read() (*data.D, error) {
	d := &data.D{}
	d.Defaults = getDefaults()

	// get cmdline data
	fset, args, err := getFlagsAndArgs()
	if err != nil {
		return d, err
	}
	d.FlagSet = fset
	d.Args = args

	// actual read of config files
	config, priv, err := readFiles(d)
	if err != nil {
		return d, err
	}
	d.Config = config
	d.Private = priv

	// put everything together on the same struct in order of precedence
	d.Config = consolidateData(d)

	return d, nil
}

// ResolveServers tries to get the server list from any source available.
// If no automatic source is configured, then it returns specified fixed list.
func ResolveServers(data *data.D) (map[string]*data.Server, error) {
	if data.Config.Gce.Project != "" {
		return gcloud.GetInstances(data)
	}
	return data.Config.Servers, nil
}
