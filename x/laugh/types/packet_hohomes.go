package types

// ValidateBasic is used for validating the packet
func (p HohomesPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p HohomesPacketData) GetBytes() ([]byte, error) {
	var modulePacket LaughPacketData

	modulePacket.Packet = &LaughPacketData_HohomesPacket{&p}

	return modulePacket.Marshal()
}
