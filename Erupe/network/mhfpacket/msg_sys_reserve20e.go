package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve20E represents the MSG_SYS_reserve20E
type MsgSysReserve20E struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve20E) Opcode() network.PacketID {
	return network.MSG_SYS_reserve20E
}

// Parse parses the packet from binary
func (m *MsgSysReserve20E) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve20E) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
