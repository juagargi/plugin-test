package plugin

import "github.com/juagargi/plugin-test/pkg/snet"

type PathHandler interface {
	Doc() string
	PathType() string
	Handler() snet.PacketHandlerFunc
}
