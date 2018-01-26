package input

// setDefaults is where you should code default values for input data
func setDefaults(d *Data) {
	// d.SetDefault("terminal", "terminator -T '{{.Title}}' -e '{{.Command}}'")
	d.SetDefault("terminal", "konsole -e \"{{.Command}}\"")
	d.SetDefault("remoteUser", "$USER")
}
