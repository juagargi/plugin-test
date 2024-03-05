package plugin

import (
	"fmt"
	"plugin"
)

type Declare func() string

func Load(filename string) *Plugin {
	// Open dynamically linked entry point.
	plugin, err := plugin.Open(filename)
	if err != nil {
		err = fmt.Errorf("loading plugin %s: %w\nType of error: %T", filename, err, err)
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
	p := &Plugin{
		Plugin:  *plugin,
		Declare: declare,
	}

	// Distribute plugin according to its declaration.
	declaration := declare()
	switch declaration {
	case "path":
		err := LoadPacketHandlerPlugin(p)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("not doing anything with %s , declared as %s\n", filename, declaration)
	}
	return p
}

type Plugin struct {
	plugin.Plugin
	Declare func() string
}

func LoadPacketHandlerPlugin(p *Plugin) error {
	return nil
}
