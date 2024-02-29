package main

import (
	"fmt"

	"github.com/juagargi/plugin-test/pkg/plugin"
)

func main() {
	fmt.Println("Running app1 ...")
	pluginFiles := plugin.List()
	for _, filepath := range pluginFiles {
		fmt.Printf("Found %s\n", filepath)
		plug := plugin.Load(filepath)
		fmt.Printf("declare says:\t\t %s\n", plug.Declare())
	}
	fmt.Println("done app1.")
}
