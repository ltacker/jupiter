package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHohosent(goCtx context.Context, msg *types.MsgCreateHohosent) (*types.MsgCreateHohosentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHohosent(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHohosentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHohosent(goCtx context.Context, msg *types.MsgUpdateHohosent) (*types.MsgUpdateHohosentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hohosent = types.Hohosent{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHohosent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHohosentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHohosent(ctx, hohosent)

	return &types.MsgUpdateHohosentResponse{}, nil
}

func (k msgServer) DeleteHohosent(goCtx context.Context, msg *types.MsgDeleteHohosent) (*types.MsgDeleteHohosentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHohosent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHohosentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHohosent(ctx, msg.Id)

	return &types.MsgDeleteHohosentResponse{}, nil
}
