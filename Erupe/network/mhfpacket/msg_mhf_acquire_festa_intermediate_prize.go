package mhfpacket

import (
	"errors"

	"erupe-ce/network"
	"erupe-ce/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireFestaIntermediatePrize represents the MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
type MsgMhfAcquireFestaIntermediatePrize struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireFestaIntermediatePrize) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireFestaIntermediatePrize) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireFestaIntermediatePrize) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
