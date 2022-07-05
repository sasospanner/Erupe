package mhfpacket

import ( 
 "errors" 

 	"erupe-ce/network/clientctx"
	"erupe-ce/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateHouse represents the MSG_MHF_UPDATE_HOUSE
type MsgMhfUpdateHouse struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateHouse) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_HOUSE
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateHouse) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateHouse) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
