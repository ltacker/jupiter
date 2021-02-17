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

func (k Keeper) HahasentAll(c context.Context, req *types.QueryAllHahasentRequest) (*types.QueryAllHahasentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hahasents []*types.Hahasent
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hahasentStore := prefix.NewStore(store, types.KeyPrefix(types.HahasentKey))

	pageRes, err := query.Paginate(hahasentStore, req.Pagination, func(key []byte, value []byte) error {
		var hahasent types.Hahasent
		if err := k.cdc.UnmarshalBinaryBare(value, &hahasent); err != nil {
			return err
		}

		hahasents = append(hahasents, &hahasent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHahasentResponse{Hahasent: hahasents, Pagination: pageRes}, nil
}

func (k Keeper) Hahasent(c context.Context, req *types.QueryGetHahasentRequest) (*types.QueryGetHahasentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hahasent types.Hahasent
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HahasentKey+req.Id)), &hahasent)

	return &types.QueryGetHahasentResponse{Hahasent: &hahasent}, nil
}
