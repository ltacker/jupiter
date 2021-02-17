package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func listHoho(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msgs := keeper.GetAllHoho(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msgs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func getHoho(ctx sdk.Context, id string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msg := keeper.GetHoho(ctx, id)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
