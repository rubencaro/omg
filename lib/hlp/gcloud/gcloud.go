// Package gcloud contains helpers to work with Google Cloud stuff
// The `gcloud` executable, actually.
package gcloud

import (
	"encoding/json"
	"fmt"

	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

// GetInstances invokes 'gcloud compute instances list --format=json'
// to get the list of instances for the given parameters
func GetInstances(data *input.Data) (map[string]*input.Server, error) {
	cmd := getInstancesCmd(data)
	res, err := hlp.Run(hlp.Silent, cmd)
	if err != nil {
		return nil, err
	}

	var f []*gceServer
	err = json.Unmarshal([]byte(res), &f)
	if err != nil {
		return nil, err
	}
	servers := parseInstances(f)
	return servers, nil
}

func getInstancesCmd(data *input.Data) string {
	return fmt.Sprintf(
		"gcloud compute instances list --format=json --project=%s",
		data.Config.Gce["project"],
	)
}

func parseInstances(f []*gceServer) map[string]*input.Server {
	servers := map[string]*input.Server{}
	for _, s := range f {
		var ip = s.NetworkInterfaces[0].AccessConfigs[0].NatIP
		servers[s.Name] = &input.Server{Name: s.Name, IP: ip}
	}
	return servers
}

type gceServer struct {
	Name              string
	NetworkInterfaces []*gceNetworkInterface
}

type gceNetworkInterface struct {
	AccessConfigs []*gceAccessConfig
}

type gceAccessConfig struct {
	NatIP string
}
