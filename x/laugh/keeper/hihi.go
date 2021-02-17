package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHihiCount get the total number of hihi
func (k Keeper) GetHihiCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiCountKey))
	byteKey := types.KeyPrefix(types.HihiCountKey)
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

// SetHihiCount set the total number of hihi
func (k Keeper) SetHihiCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiCountKey))
	byteKey := types.KeyPrefix(types.HihiCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHihi appends a hihi in the store with a new id and update the count
func (k Keeper) AppendHihi(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the hihi
	count := k.GetHihiCount(ctx)
	id := strconv.FormatInt(count, 10)
	var hihi = types.Hihi{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	key := types.KeyPrefix(types.HihiKey + hihi.Id)
	value := k.cdc.MustMarshalBinaryBare(&hihi)
	store.Set(key, value)

	// Update hihi count
	k.SetHihiCount(ctx, count+1)

	return id
}

// SetHihi set a specific hihi in the store
func (k Keeper) SetHihi(ctx sdk.Context, hihi types.Hihi) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	b := k.cdc.MustMarshalBinaryBare(&hihi)
	store.Set(types.KeyPrefix(types.HihiKey+hihi.Id), b)
}

// GetHihi returns a hihi from its id
func (k Keeper) GetHihi(ctx sdk.Context, key string) types.Hihi {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	var hihi types.Hihi
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HihiKey+key)), &hihi)
	return hihi
}

// HasHihi checks if the hihi exists in the store
func (k Keeper) HasHihi(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	return store.Has(types.KeyPrefix(types.HihiKey + id))
}

// GetHihiOwner returns the creator of the hihi
func (k Keeper) GetHihiOwner(ctx sdk.Context, key string) string {
	return k.GetHihi(ctx, key).Creator
}

// DeleteHihi removes a hihi from the store
func (k Keeper) RemoveHihi(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	store.Delete(types.KeyPrefix(types.HihiKey + key))
}

// GetAllHihi returns all hihi
func (k Keeper) GetAllHihi(ctx sdk.Context) (msgs []types.Hihi) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HihiKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HihiKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Hihi
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
