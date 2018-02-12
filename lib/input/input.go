package input

import (
	"regexp"
	"strings"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/input/flags"
	"github.com/rubencaro/omg/lib/input/gcloud"
)

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
func ResolveServers(d *data.D) (map[string]*data.Server, error) {
	complete, err := getCompleteServerList(d)
	if err != nil {
		return nil, err
	}
	filtered := applyServersFlag(complete, d)
	return applyMatch(filtered, d)
}

func applyMatch(in map[string]*data.Server, d *data.D) (map[string]*data.Server, error) {
	match := getMatchString(d)
	if match == "" {
		return in, nil
	}
	rgx, err := regexp.Compile(match)
	if err != nil {
		return nil, err
	}
	out := map[string]*data.Server{}
	for k := range in {
		if rgx.MatchString(in[k].Name) {
			out[k] = in[k]
		}
	}
	return out, nil
}

func getMatchString(d *data.D) string {
	match, _ := flags.GetFlag(d, "match")
	if match == "" {
		match = d.Config.Gce.Match
	}
	return match
}

func applyServersFlag(in map[string]*data.Server, d *data.D) map[string]*data.Server {
	fixed, _ := flags.GetFlag(d, "servers")
	if fixed != "" {
		names := strings.Split(fixed, ",")
		return mapNamesToServers(names, in)
	}
	return in
}

func getCompleteServerList(d *data.D) (map[string]*data.Server, error) {
	if d.Config.Gce.Project != "" {
		return gcloud.GetInstances(d)
	}
	return d.Config.Servers, nil
}

func mapNamesToServers(names []string, servers map[string]*data.Server) map[string]*data.Server {
	res := map[string]*data.Server{}
	for i := range names {
		res[names[i]] = servers[names[i]]
	}
	return res
}
