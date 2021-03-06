package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/ltacker/jupiter/x/laugh/types"
)

// TransmitHahamesPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitHahamesPacket(
	ctx sdk.Context,
	packetData types.HahamesPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvHahamesPacket processes packet reception
func (k Keeper) OnRecvHahamesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HahamesPacketData) error {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return err
	}

	// TODO: packet reception logic
	k.AppendHaha(ctx, packet.SourcePort+"-"+packet.SourceChannel, "")

	return nil
}

// OnAcknowledgementHahamesPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementHahamesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HahamesPacketData, ack channeltypes.Acknowledgement) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		k.AppendHahasent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "fail")

		return nil
	default:

		// TODO: successful acknowledgement logic
		k.AppendHahasent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "success")

		return nil
	}
}

// OnTimeoutHahamesPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutHahamesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HahamesPacketData) error {

	// TODO: packet timeout logic
	k.AppendHahasent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "timeout")

	return nil
}
