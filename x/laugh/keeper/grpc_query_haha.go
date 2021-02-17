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

func (k Keeper) HahaAll(c context.Context, req *types.QueryAllHahaRequest) (*types.QueryAllHahaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hahas []*types.Haha
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hahaStore := prefix.NewStore(store, types.KeyPrefix(types.HahaKey))

	pageRes, err := query.Paginate(hahaStore, req.Pagination, func(key []byte, value []byte) error {
		var haha types.Haha
		if err := k.cdc.UnmarshalBinaryBare(value, &haha); err != nil {
			return err
		}

		hahas = append(hahas, &haha)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHahaResponse{Haha: hahas, Pagination: pageRes}, nil
}

func (k Keeper) Haha(c context.Context, req *types.QueryGetHahaRequest) (*types.QueryGetHahaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var haha types.Haha
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HahaKey+req.Id)), &haha)

	return &types.QueryGetHahaResponse{Haha: &haha}, nil
}
