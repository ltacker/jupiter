package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/ltacker/jupiter/x/laugh/types"
)

// TransmitHohomesPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitHohomesPacket(
	ctx sdk.Context,
	packetData types.HohomesPacketData,
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

// OnRecvHohomesPacket processes packet reception
func (k Keeper) OnRecvHohomesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HohomesPacketData) error {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return err
	}

	// TODO: packet reception logic
	k.AppendHoho(ctx, packet.SourcePort+"-"+packet.SourceChannel, "")

	return nil
}

// OnAcknowledgementHohomesPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementHohomesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HohomesPacketData, ack channeltypes.Acknowledgement) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		k.AppendHohosent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "fail")

		return nil
	default:

		// TODO: successful acknowledgement logic
		k.AppendHohosent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "success")

		return nil
	}
}

// OnTimeoutHohomesPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutHohomesPacket(ctx sdk.Context, packet channeltypes.Packet, data types.HohomesPacketData) error {

	// TODO: packet timeout logic
	k.AppendHohosent(ctx, packet.SourcePort+"-"+packet.SourceChannel, "timeout")

	return nil
}
