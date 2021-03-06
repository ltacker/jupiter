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

func (k Keeper) HihisentAll(c context.Context, req *types.QueryAllHihisentRequest) (*types.QueryAllHihisentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hihisents []*types.Hihisent
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hihisentStore := prefix.NewStore(store, types.KeyPrefix(types.HihisentKey))

	pageRes, err := query.Paginate(hihisentStore, req.Pagination, func(key []byte, value []byte) error {
		var hihisent types.Hihisent
		if err := k.cdc.UnmarshalBinaryBare(value, &hihisent); err != nil {
			return err
		}

		hihisents = append(hihisents, &hihisent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHihisentResponse{Hihisent: hihisents, Pagination: pageRes}, nil
}

func (k Keeper) Hihisent(c context.Context, req *types.QueryGetHihisentRequest) (*types.QueryGetHihisentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hihisent types.Hihisent
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HihisentKey+req.Id)), &hihisent)

	return &types.QueryGetHihisentResponse{Hihisent: &hihisent}, nil
}
