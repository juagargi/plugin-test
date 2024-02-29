package snet

import "fmt"

// Global registry of packet handlers.
var PacketHandlers map[string]PacketHandlerFunc

// Handler function.
type PacketHandlerFunc func(p *Packet) error

func AddHandler(packetType string, handler PacketHandlerFunc) error {
	if _, ok := PacketHandlers[packetType]; ok {
		return fmt.Errorf("there already exists a handler for this type: %s", packetType)
	}
	PacketHandlers[packetType] = handler
	return nil
}

// Packet is a packet.
type Packet struct {
	TypeOfPacket string
	SizeOfPacket int
	Ingress      uint16
	Egress       uint16
	CurrentHop   int
}

func (p *Packet) String() string {
	return fmt.Sprintf("snet.Packet type: %s , size: %d", p.TypeOfPacket, p.SizeOfPacket)
}
