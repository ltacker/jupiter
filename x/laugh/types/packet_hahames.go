package types

// ValidateBasic is used for validating the packet
func (p HahamesPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p HahamesPacketData) GetBytes() ([]byte, error) {
	var modulePacket LaughPacketData

	modulePacket.Packet = &LaughPacketData_HahamesPacket{&p}

	return modulePacket.Marshal()
}
