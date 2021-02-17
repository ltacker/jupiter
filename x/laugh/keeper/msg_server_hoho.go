package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHoho(goCtx context.Context, msg *types.MsgCreateHoho) (*types.MsgCreateHohoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHoho(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHohoResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHoho(goCtx context.Context, msg *types.MsgUpdateHoho) (*types.MsgUpdateHohoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hoho = types.Hoho{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHoho(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHohoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHoho(ctx, hoho)

	return &types.MsgUpdateHohoResponse{}, nil
}

func (k msgServer) DeleteHoho(goCtx context.Context, msg *types.MsgDeleteHoho) (*types.MsgDeleteHohoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHoho(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHohoOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHoho(ctx, msg.Id)

	return &types.MsgDeleteHohoResponse{}, nil
}
