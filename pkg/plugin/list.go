package plugin

import "path/filepath"

func List() []string {
	// Get plugins from the current directory.
	p, err := filepath.Glob("./bin/plugins/*.so")
	if err != nil {
		panic(err)
	}
	return p
}
