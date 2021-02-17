package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) SendHohomes(goCtx context.Context, msg *types.MsgSendHohomes) (*types.MsgSendHohomesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.HohomesPacketData

	// Transmit the packet
	err := k.TransmitHohomesPacket(
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

	return &types.MsgSendHohomesResponse{}, nil
}
