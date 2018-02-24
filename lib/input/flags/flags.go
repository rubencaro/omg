// Package flags contains helpers to get values out of flags
package flags

import (
	"os"

	"github.com/rubencaro/omg/lib/data"
)

// GetFlagOrEnv looks for given key on flags, and if it doesn't find a value,
// then it looks for a 'OMG_<key>' environment variable.
// Returns empty string if not found.
func GetFlagOrEnv(d *data.D, key string) string {
	v, ok := GetFlag(d, key)
	if ok {
		return v
	}
	return os.Getenv("OMG_" + key)
}

// GetFlag returns the string value of a flag contained in given D,
// and a bool indicating whether it was found
func GetFlag(d *data.D, key string) (string, bool) {
	f := d.FlagSet.Lookup(key)
	if f == nil {
		return "", false
	}
	return f.Value.String(), true
}

// GetBoolFlag returns the bool value of a flag contained in given D,
// and a bool indicating whether it was found
func GetBoolFlag(d *data.D, key string) (bool, bool) {
	f := d.FlagSet.Lookup(key)
	if f == nil {
		return false, false
	}
	return f.Value.String() == "true", true
}
