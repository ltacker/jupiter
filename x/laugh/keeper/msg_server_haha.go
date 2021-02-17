package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHaha(goCtx context.Context, msg *types.MsgCreateHaha) (*types.MsgCreateHahaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHaha(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHahaResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHaha(goCtx context.Context, msg *types.MsgUpdateHaha) (*types.MsgUpdateHahaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var haha = types.Haha{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHaha(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHahaOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHaha(ctx, haha)

	return &types.MsgUpdateHahaResponse{}, nil
}

func (k msgServer) DeleteHaha(goCtx context.Context, msg *types.MsgDeleteHaha) (*types.MsgDeleteHahaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHaha(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHahaOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHaha(ctx, msg.Id)

	return &types.MsgDeleteHahaResponse{}, nil
}
