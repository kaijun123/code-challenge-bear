package keeper

import (
	"encoding/binary"
	"fmt"

	types "bear/x/bear/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBearIDBytes
func GetBearIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBearCount
func (k Keeper) GetBearCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BearCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// SetBearCount
func (k Keeper) SetBearCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BearCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBear
func (k Keeper) AppendBear(ctx sdk.Context, bear types.Bear) uint64 {
	count := k.GetBearCount(ctx)
	bear.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BearKey))
	appendedValue := k.cdc.MustMarshal(&bear)
	store.Set(GetBearIDBytes(bear.Id), appendedValue)
	k.SetBearCount(ctx, count+1)
	return count
}

// GetBear
func (k Keeper) GetBear(ctx sdk.Context, id uint64) (val types.Bear, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BearKey))
	b := store.Get(GetBearIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetBear
func (k Keeper) SetBear(ctx sdk.Context, bear types.Bear) {
	fmt.Println("id set: ", bear.Id)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BearKey))
	b := k.cdc.MustMarshal(&bear)
	store.Set(GetBearIDBytes(bear.Id), b)
}

// RemoveBear
func (k Keeper) RemoveBear(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BearKey))
	store.Delete(GetBearIDBytes(id))
}
