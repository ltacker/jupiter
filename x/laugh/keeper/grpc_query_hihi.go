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

func (k Keeper) HihiAll(c context.Context, req *types.QueryAllHihiRequest) (*types.QueryAllHihiResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hihis []*types.Hihi
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hihiStore := prefix.NewStore(store, types.KeyPrefix(types.HihiKey))

	pageRes, err := query.Paginate(hihiStore, req.Pagination, func(key []byte, value []byte) error {
		var hihi types.Hihi
		if err := k.cdc.UnmarshalBinaryBare(value, &hihi); err != nil {
			return err
		}

		hihis = append(hihis, &hihi)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHihiResponse{Hihi: hihis, Pagination: pageRes}, nil
}

func (k Keeper) Hihi(c context.Context, req *types.QueryGetHihiRequest) (*types.QueryGetHihiResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hihi types.Hihi
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HihiKey+req.Id)), &hihi)

	return &types.QueryGetHihiResponse{Hihi: &hihi}, nil
}
