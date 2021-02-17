package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ltacker/jupiter/x/laugh/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HohosentAll(c context.Context, req *types.QueryAllHohosentRequest) (*types.QueryAllHohosentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hohosents []*types.Hohosent
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hohosentStore := prefix.NewStore(store, types.KeyPrefix(types.HohosentKey))

	pageRes, err := query.Paginate(hohosentStore, req.Pagination, func(key []byte, value []byte) error {
		var hohosent types.Hohosent
		if err := k.cdc.UnmarshalBinaryBare(value, &hohosent); err != nil {
			return err
		}

		hohosents = append(hohosents, &hohosent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHohosentResponse{Hohosent: hohosents, Pagination: pageRes}, nil
}

func (k Keeper) Hohosent(c context.Context, req *types.QueryGetHohosentRequest) (*types.QueryGetHohosentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hohosent types.Hohosent
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HohosentKey+req.Id)), &hohosent)

	return &types.QueryGetHohosentResponse{Hohosent: &hohosent}, nil
}
