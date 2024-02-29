package plugin

import (
	"fmt"
	"plugin"
)

func Load(filename string) *Plugin {
	// Open dynamically linked entry point.
	plugin, err := plugin.Open(filename)
	if err != nil {
		err = fmt.Errorf("loading plugin %s: %w", filename, err)
		panic(err)
	}
	// Check for generic Declare symbol.
	symb, err := plugin.Lookup("Declare")
	if err != nil {
		err = fmt.Errorf("checking plugin type %s: %w", filename, err)
		panic(err)
	}
	// Is it the right type?
	declare, ok := symb.(func() string)
	if !ok {
		panic("bad signature")
	}

	// Wrap the plugin.
	p := Plugin{
		Plugin:  *plugin,
		Declare: declare,
	}
	return &p
}

type Plugin struct {
	plugin.Plugin
	Declare func() string
}
