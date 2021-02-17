package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/jupiter/x/laugh/types"
)

func (k msgServer) CreateHihisent(goCtx context.Context, msg *types.MsgCreateHihisent) (*types.MsgCreateHihisentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendHihisent(
		ctx,
		msg.Creator,
		msg.Text,
	)

	return &types.MsgCreateHihisentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHihisent(goCtx context.Context, msg *types.MsgUpdateHihisent) (*types.MsgUpdateHihisentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hihisent = types.Hihisent{
		Creator: msg.Creator,
		Id:      msg.Id,
		Text:    msg.Text,
	}

	// Checks that the element exists
	if !k.HasHihisent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHihisentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHihisent(ctx, hihisent)

	return &types.MsgUpdateHihisentResponse{}, nil
}

func (k msgServer) DeleteHihisent(goCtx context.Context, msg *types.MsgDeleteHihisent) (*types.MsgDeleteHihisentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHihisent(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHihisentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHihisent(ctx, msg.Id)

	return &types.MsgDeleteHihisentResponse{}, nil
}
