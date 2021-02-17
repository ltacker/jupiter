package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHohosentCount get the total number of hohosent
func (k Keeper) GetHohosentCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentCountKey))
	byteKey := types.KeyPrefix(types.HohosentCountKey)
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

// SetHohosentCount set the total number of hohosent
func (k Keeper) SetHohosentCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentCountKey))
	byteKey := types.KeyPrefix(types.HohosentCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHohosent appends a hohosent in the store with a new id and update the count
func (k Keeper) AppendHohosent(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the hohosent
	count := k.GetHohosentCount(ctx)
	id := strconv.FormatInt(count, 10)
	var hohosent = types.Hohosent{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	key := types.KeyPrefix(types.HohosentKey + hohosent.Id)
	value := k.cdc.MustMarshalBinaryBare(&hohosent)
	store.Set(key, value)

	// Update hohosent count
	k.SetHohosentCount(ctx, count+1)

	return id
}

// SetHohosent set a specific hohosent in the store
func (k Keeper) SetHohosent(ctx sdk.Context, hohosent types.Hohosent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	b := k.cdc.MustMarshalBinaryBare(&hohosent)
	store.Set(types.KeyPrefix(types.HohosentKey+hohosent.Id), b)
}

// GetHohosent returns a hohosent from its id
func (k Keeper) GetHohosent(ctx sdk.Context, key string) types.Hohosent {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	var hohosent types.Hohosent
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HohosentKey+key)), &hohosent)
	return hohosent
}

// HasHohosent checks if the hohosent exists in the store
func (k Keeper) HasHohosent(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	return store.Has(types.KeyPrefix(types.HohosentKey + id))
}

// GetHohosentOwner returns the creator of the hohosent
func (k Keeper) GetHohosentOwner(ctx sdk.Context, key string) string {
	return k.GetHohosent(ctx, key).Creator
}

// DeleteHohosent removes a hohosent from the store
func (k Keeper) RemoveHohosent(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	store.Delete(types.KeyPrefix(types.HohosentKey + key))
}

// GetAllHohosent returns all hohosent
func (k Keeper) GetAllHohosent(ctx sdk.Context) (msgs []types.Hohosent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohosentKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HohosentKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Hohosent
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
