package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHihi(goCtx context.Context, msg *types.MsgCreateHihi) (*types.MsgCreateHihiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHihi(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHihiResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHihi(goCtx context.Context, msg *types.MsgUpdateHihi) (*types.MsgUpdateHihiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hihi = types.Hihi{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHihi(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHihiOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHihi(ctx, hihi)

	return &types.MsgUpdateHihiResponse{}, nil
}

func (k msgServer) DeleteHihi(goCtx context.Context, msg *types.MsgDeleteHihi) (*types.MsgDeleteHihiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHihi(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHihiOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHihi(ctx, msg.Id)

	return &types.MsgDeleteHihiResponse{}, nil
}
