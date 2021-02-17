package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHihisentCount get the total number of hihisent
func (k Keeper) GetHihisentCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentCountKey))
	byteKey := types.KeyPrefix(types.HihisentCountKey)
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

// SetHihisentCount set the total number of hihisent
func (k Keeper) SetHihisentCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentCountKey))
	byteKey := types.KeyPrefix(types.HihisentCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHihisent appends a hihisent in the store with a new id and update the count
func (k Keeper) AppendHihisent(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the hihisent
	count := k.GetHihisentCount(ctx)
	id := strconv.FormatInt(count, 10)
	var hihisent = types.Hihisent{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	key := types.KeyPrefix(types.HihisentKey + hihisent.Id)
	value := k.cdc.MustMarshalBinaryBare(&hihisent)
	store.Set(key, value)

	// Update hihisent count
	k.SetHihisentCount(ctx, count+1)

	return id
}

// SetHihisent set a specific hihisent in the store
func (k Keeper) SetHihisent(ctx sdk.Context, hihisent types.Hihisent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	b := k.cdc.MustMarshalBinaryBare(&hihisent)
	store.Set(types.KeyPrefix(types.HihisentKey+hihisent.Id), b)
}

// GetHihisent returns a hihisent from its id
func (k Keeper) GetHihisent(ctx sdk.Context, key string) types.Hihisent {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	var hihisent types.Hihisent
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HihisentKey+key)), &hihisent)
	return hihisent
}

// HasHihisent checks if the hihisent exists in the store
func (k Keeper) HasHihisent(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	return store.Has(types.KeyPrefix(types.HihisentKey + id))
}

// GetHihisentOwner returns the creator of the hihisent
func (k Keeper) GetHihisentOwner(ctx sdk.Context, key string) string {
	return k.GetHihisent(ctx, key).Creator
}

// DeleteHihisent removes a hihisent from the store
func (k Keeper) RemoveHihisent(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	store.Delete(types.KeyPrefix(types.HihisentKey + key))
}

// GetAllHihisent returns all hihisent
func (k Keeper) GetAllHihisent(ctx sdk.Context) (msgs []types.Hihisent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihisentKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HihisentKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Hihisent
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
