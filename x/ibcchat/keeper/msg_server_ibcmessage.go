package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/ltacker/jupiter/x/ibcchat/types"
)

func (k msgServer) SendIbcmessage(goCtx context.Context, msg *types.MsgSendIbcmessage) (*types.MsgSendIbcmessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcmessagePacketData

	packet.Text = msg.Text

	// Transmit the packet
	err := k.TransmitIbcmessagePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcmessageResponse{}, nil
}
