package input

import "github.com/rubencaro/omg/lib/data"

// consolidateData gets the value for each possible configuration option
// and returns them on a data.Config struct.
// Every one of them may be calculated in a different way, so it's cleaner
// just to code them one by one than code a generic merge of structs and then deal
// with a lot of corner cases.
func consolidateData(d *data.D) *data.Config {
	return &data.Config{
		Terminal:   getTerminal(d),
		RemoteUser: getRemoteUser(d),
		Customs:    getCustoms(d),
		Servers:    getServers(d),
		Gce:        getGce(d),
	}
}

func getTerminal(d *data.D) string {
	return getFNZString(d.Private.Terminal, d.Config.Terminal, d.Defaults.Terminal)
}

func getRemoteUser(d *data.D) string {
	return getFNZString(d.Private.RemoteUser, d.Config.RemoteUser, d.Defaults.RemoteUser)
}

func getCustoms(d *data.D) map[string]*data.Custom {
	return getFNZMapStringCustom(d.Private.Customs, d.Config.Customs)
}

func getServers(d *data.D) map[string]*data.Server {
	s := d.Config.Servers
	if len(d.Private.Servers) > 0 {
		s = d.Private.Servers
	}
	formatFixedServers(s)
	return s
}

func formatFixedServers(servers map[string]*data.Server) {
	for k, v := range servers {
		v.Name = k
	}
}

func getGce(d *data.D) *data.Gce {
	if d.Private.Gce != nil && d.Private.Gce.Project != "" {
		return d.Private.Gce
	}
	return d.Config.Gce
}
