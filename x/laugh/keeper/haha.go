package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHahaCount get the total number of haha
func (k Keeper) GetHahaCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaCountKey))
	byteKey := types.KeyPrefix(types.HahaCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetHahaCount set the total number of haha
func (k Keeper) SetHahaCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaCountKey))
	byteKey := types.KeyPrefix(types.HahaCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHaha appends a haha in the store with a new id and update the count
func (k Keeper) AppendHaha(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the haha
	count := k.GetHahaCount(ctx)
	id := strconv.FormatInt(count, 10)
	var haha = types.Haha{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	key := types.KeyPrefix(types.HahaKey + haha.Id)
	value := k.cdc.MustMarshalBinaryBare(&haha)
	store.Set(key, value)

	// Update haha count
	k.SetHahaCount(ctx, count+1)

	return id
}

// SetHaha set a specific haha in the store
func (k Keeper) SetHaha(ctx sdk.Context, haha types.Haha) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	b := k.cdc.MustMarshalBinaryBare(&haha)
	store.Set(types.KeyPrefix(types.HahaKey+haha.Id), b)
}

// GetHaha returns a haha from its id
func (k Keeper) GetHaha(ctx sdk.Context, key string) types.Haha {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	var haha types.Haha
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HahaKey+key)), &haha)
	return haha
}

// HasHaha checks if the haha exists in the store
func (k Keeper) HasHaha(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	return store.Has(types.KeyPrefix(types.HahaKey + id))
}

// GetHahaOwner returns the creator of the haha
func (k Keeper) GetHahaOwner(ctx sdk.Context, key string) string {
	return k.GetHaha(ctx, key).Creator
}

// DeleteHaha removes a haha from the store
func (k Keeper) RemoveHaha(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	store.Delete(types.KeyPrefix(types.HahaKey + key))
}

// GetAllHaha returns all haha
func (k Keeper) GetAllHaha(ctx sdk.Context) (msgs []types.Haha) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahaKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HahaKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Haha
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
