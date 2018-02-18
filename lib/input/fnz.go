package input

// FNZ (FirstNonZero) functions

import "github.com/rubencaro/omg/lib/data"

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

func getFNZMapStringCustom(values ...map[string]*data.Custom) map[string]*data.Custom {
	for _, v := range values {
		if len(v) > 0 {
			return v
		}
	}
	return map[string]*data.Custom{}
}
