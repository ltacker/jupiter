package types

// ValidateBasic is used for validating the packet
func (p IbcmessagePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcmessagePacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcchatPacketData

	modulePacket.Packet = &IbcchatPacketData_IbcmessagePacket{&p}

	return modulePacket.Marshal()
}
