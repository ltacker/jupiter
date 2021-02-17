package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHahasentCount get the total number of hahasent
func (k Keeper) GetHahasentCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentCountKey))
	byteKey := types.KeyPrefix(types.HahasentCountKey)
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

// SetHahasentCount set the total number of hahasent
func (k Keeper) SetHahasentCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentCountKey))
	byteKey := types.KeyPrefix(types.HahasentCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHahasent appends a hahasent in the store with a new id and update the count
func (k Keeper) AppendHahasent(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the hahasent
	count := k.GetHahasentCount(ctx)
	id := strconv.FormatInt(count, 10)
	var hahasent = types.Hahasent{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	key := types.KeyPrefix(types.HahasentKey + hahasent.Id)
	value := k.cdc.MustMarshalBinaryBare(&hahasent)
	store.Set(key, value)

	// Update hahasent count
	k.SetHahasentCount(ctx, count+1)

	return id
}

// SetHahasent set a specific hahasent in the store
func (k Keeper) SetHahasent(ctx sdk.Context, hahasent types.Hahasent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	b := k.cdc.MustMarshalBinaryBare(&hahasent)
	store.Set(types.KeyPrefix(types.HahasentKey+hahasent.Id), b)
}

// GetHahasent returns a hahasent from its id
func (k Keeper) GetHahasent(ctx sdk.Context, key string) types.Hahasent {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	var hahasent types.Hahasent
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HahasentKey+key)), &hahasent)
	return hahasent
}

// HasHahasent checks if the hahasent exists in the store
func (k Keeper) HasHahasent(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	return store.Has(types.KeyPrefix(types.HahasentKey + id))
}

// GetHahasentOwner returns the creator of the hahasent
func (k Keeper) GetHahasentOwner(ctx sdk.Context, key string) string {
	return k.GetHahasent(ctx, key).Creator
}

// DeleteHahasent removes a hahasent from the store
func (k Keeper) RemoveHahasent(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	store.Delete(types.KeyPrefix(types.HahasentKey + key))
}

// GetAllHahasent returns all hahasent
func (k Keeper) GetAllHahasent(ctx sdk.Context) (msgs []types.Hahasent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HahasentKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HahasentKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Hahasent
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
