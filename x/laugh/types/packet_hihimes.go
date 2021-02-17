package types

// ValidateBasic is used for validating the packet
func (p HihimesPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p HihimesPacketData) GetBytes() ([]byte, error) {
	var modulePacket LaughPacketData

	modulePacket.Packet = &LaughPacketData_HihimesPacket{&p}

	return modulePacket.Marshal()
}
