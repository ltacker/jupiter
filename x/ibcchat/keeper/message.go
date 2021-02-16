package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/ibcchat/types"
	"strconv"
)

// GetMessageCount get the total number of message
func (k Keeper) GetMessageCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageCountKey))
	byteKey := types.KeyPrefix(types.MessageCountKey)
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

// SetMessageCount set the total number of message
func (k Keeper) SetMessageCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageCountKey))
	byteKey := types.KeyPrefix(types.MessageCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendMessage appends a message in the store with a new id and update the count
func (k Keeper) AppendMessage(
	ctx sdk.Context,
	creator string,
	text string,
) string {
	// Create the message
	count := k.GetMessageCount(ctx)
	id := strconv.FormatInt(count, 10)
	var message = types.Message{
		Creator: creator,
		Id:      id,
		Text:    text,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	key := types.KeyPrefix(types.MessageKey + message.Id)
	value := k.cdc.MustMarshalBinaryBare(&message)
	store.Set(key, value)

	// Update message count
	k.SetMessageCount(ctx, count+1)

	return id
}

// SetMessage set a specific message in the store
func (k Keeper) SetMessage(ctx sdk.Context, message types.Message) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	b := k.cdc.MustMarshalBinaryBare(&message)
	store.Set(types.KeyPrefix(types.MessageKey+message.Id), b)
}

// GetMessage returns a message from its id
func (k Keeper) GetMessage(ctx sdk.Context, key string) types.Message {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	var message types.Message
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.MessageKey+key)), &message)
	return message
}

// HasMessage checks if the message exists in the store
func (k Keeper) HasMessage(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	return store.Has(types.KeyPrefix(types.MessageKey + id))
}

// GetMessageOwner returns the creator of the message
func (k Keeper) GetMessageOwner(ctx sdk.Context, key string) string {
	return k.GetMessage(ctx, key).Creator
}

// DeleteMessage removes a message from the store
func (k Keeper) RemoveMessage(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	store.Delete(types.KeyPrefix(types.MessageKey + key))
}

// GetAllMessage returns all message
func (k Keeper) GetAllMessage(ctx sdk.Context) (msgs []types.Message) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.MessageKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Message
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
