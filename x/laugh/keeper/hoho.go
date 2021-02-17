package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/types"
	"strconv"
)

// GetHohoCount get the total number of hoho
func (k Keeper) GetHohoCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoCountKey))
	byteKey := types.KeyPrefix(types.HohoCountKey)
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

// SetHohoCount set the total number of hoho
func (k Keeper) SetHohoCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoCountKey))
	byteKey := types.KeyPrefix(types.HohoCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendHoho appends a hoho in the store with a new id and update the count
func (k Keeper) AppendHoho(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the hoho
	count := k.GetHohoCount(ctx)
	id := strconv.FormatInt(count, 10)
	var hoho = types.Hoho{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	key := types.KeyPrefix(types.HohoKey + hoho.Id)
	value := k.cdc.MustMarshalBinaryBare(&hoho)
	store.Set(key, value)

	// Update hoho count
	k.SetHohoCount(ctx, count+1)

	return id
}

// SetHoho set a specific hoho in the store
func (k Keeper) SetHoho(ctx sdk.Context, hoho types.Hoho) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	b := k.cdc.MustMarshalBinaryBare(&hoho)
	store.Set(types.KeyPrefix(types.HohoKey+hoho.Id), b)
}

// GetHoho returns a hoho from its id
func (k Keeper) GetHoho(ctx sdk.Context, key string) types.Hoho {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	var hoho types.Hoho
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.HohoKey+key)), &hoho)
	return hoho
}

// HasHoho checks if the hoho exists in the store
func (k Keeper) HasHoho(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	return store.Has(types.KeyPrefix(types.HohoKey + id))
}

// GetHohoOwner returns the creator of the hoho
func (k Keeper) GetHohoOwner(ctx sdk.Context, key string) string {
	return k.GetHoho(ctx, key).Creator
}

// DeleteHoho removes a hoho from the store
func (k Keeper) RemoveHoho(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	store.Delete(types.KeyPrefix(types.HohoKey + key))
}

// GetAllHoho returns all hoho
func (k Keeper) GetAllHoho(ctx sdk.Context) (msgs []types.Hoho) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HohoKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.HohoKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Hoho
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
