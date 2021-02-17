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

func (k Keeper) HohoAll(c context.Context, req *types.QueryAllHohoRequest) (*types.QueryAllHohoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hohos []*types.Hoho
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hohoStore := prefix.NewStore(store, types.KeyPrefix(types.HohoKey))

	pageRes, err := query.Paginate(hohoStore, req.Pagination, func(key []byte, value []byte) error {
		var hoho types.Hoho
		if err := k.cdc.UnmarshalBinaryBare(value, &hoho); err != nil {
			return err
		}

		hohos = append(hohos, &hoho)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHohoResponse{Hoho: hohos, Pagination: pageRes}, nil
}

func (k Keeper) Hoho(c context.Context, req *types.QueryGetHohoRequest) (*types.QueryGetHohoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hoho types.Hoho
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HohoKey+req.Id)), &hoho)

	return &types.QueryGetHohoResponse{Hoho: &hoho}, nil
}
