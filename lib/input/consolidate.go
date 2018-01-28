package input

// consolidateData gets the value for each possible configuration option
// and returns them on a ConfigData struct.
// Every one of them may be calculated in a different way, so it's cleaner
// just to code them one by one than code a generic merge of structs and then deal
// with a lot of corner cases.
func consolidateData(d *Data) *ConfigData {
	return &ConfigData{
		Terminal:   getTerminal(d),
		RemoteUser: getRemoteUser(d),
		Custom:     getCustom(d),
		Servers:    getServers(d),
		Gce:        getGce(d),
	}
}

func getTerminal(d *Data) string {
	return getFNZString(d.Private.Terminal, d.Config.Terminal, d.Defaults.Terminal)
}

func getRemoteUser(d *Data) string {
	return getFNZString(d.Private.RemoteUser, d.Config.RemoteUser, d.Defaults.RemoteUser)
}

func getCustom(d *Data) map[string]string {
	return getFNZMapStringString(d.Private.Custom, d.Config.Custom)
}

func getServers(d *Data) map[string]*Server {
	s := d.Config.Servers
	if len(d.Private.Servers) > 0 {
		s = d.Private.Servers
	}
	formatFixedServers(s)
	return s
}

func formatFixedServers(servers map[string]*Server) {
	for k, v := range servers {
		v.Name = k
	}
}

func getGce(d *Data) map[string]string {
	return getFNZMapStringString(d.Private.Gce, d.Config.Gce)
}

// getF is a convenience for GetFlagOrEnv
func getF(d *Data, key string) string {
	return GetFlagOrEnv(d, key)
}

// FNZ (FirstNonZero) functions

func getFNZString(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func getFNZMapStringString(values ...map[string]string) map[string]string {
	for _, v := range values {
		if len(v) > 0 {
			return v
		}
	}
	return map[string]string{}
}
