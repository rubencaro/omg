// Package gcloud contains helpers to work with Google Cloud stuff
// The `gcloud` executable, actually.
package gcloud

import (
	"encoding/json"
	"fmt"

	"github.com/rubencaro/omg/lib/data"
	"github.com/rubencaro/omg/lib/hlp"
)

// GetInstances invokes 'gcloud compute instances list --format=json'
// to get the list of instances for the given parameters
func GetInstances(data *data.D) (map[string]*data.Server, error) {
	cmd := getInstancesCmd(data)
	res, err := hlp.Run(cmd, nil)
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

func getInstancesCmd(data *data.D) string {
	return fmt.Sprintf(
		"gcloud compute instances list --format=json --project=%s",
		data.Config.Gce.Project,
	)
}

func parseInstances(f []*gceServer) map[string]*data.Server {
	servers := map[string]*data.Server{}
	for _, s := range f {
		var ip = s.NetworkInterfaces[0].AccessConfigs[0].NatIP
		servers[s.Name] = &data.Server{Name: s.Name, IP: ip}
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
