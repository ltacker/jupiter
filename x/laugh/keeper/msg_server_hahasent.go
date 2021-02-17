package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHahasent(goCtx context.Context, msg *types.MsgCreateHahasent) (*types.MsgCreateHahasentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHahasent(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHahasentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHahasent(goCtx context.Context, msg *types.MsgUpdateHahasent) (*types.MsgUpdateHahasentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hahasent = types.Hahasent{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHahasent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHahasentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHahasent(ctx, hahasent)

	return &types.MsgUpdateHahasentResponse{}, nil
}

func (k msgServer) DeleteHahasent(goCtx context.Context, msg *types.MsgDeleteHahasent) (*types.MsgDeleteHahasentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHahasent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHahasentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHahasent(ctx, msg.Id)

	return &types.MsgDeleteHahasentResponse{}, nil
}
