// Package gcloud contains helpers to work with Google Cloud stuff
// The `gcloud` executable, actually.
package gcloud

import (
	"github.com/rubencaro/omg/lib/hlp"
	"github.com/rubencaro/omg/lib/input"
)

// Instances invokes 'gcloud compute instances list --format=json'
// to get the list of instances for the given parameters
func Instances(d *input.Data) (string, error) {
	// TODO https://github.com/rubencaro/bottler/blob/master/lib/bottler/helpers/gce.ex
	return hlp.Run("echo hey")
}
