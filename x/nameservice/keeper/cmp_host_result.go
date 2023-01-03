package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// SetCmpHostResult set a specific cmpHostResult in the store from its index
func (k Keeper) SetCmpHostResult(ctx sdk.Context, cmpHostResult types.CmpHostResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpHostResultKeyPrefix))
	b := k.cdc.MustMarshal(&cmpHostResult)
	store.Set(types.CmpHostResultKey(
		cmpHostResult.Index,
	), b)
}

// GetCmpHostResult returns a cmpHostResult from its index
func (k Keeper) GetCmpHostResult(
	ctx sdk.Context,
	index string,

) (val types.CmpHostResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpHostResultKeyPrefix))

	b := store.Get(types.CmpHostResultKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCmpHostResult removes a cmpHostResult from the store
func (k Keeper) RemoveCmpHostResult(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpHostResultKeyPrefix))
	store.Delete(types.CmpHostResultKey(
		index,
	))
}

// GetAllCmpHostResult returns all cmpHostResult
func (k Keeper) GetAllCmpHostResult(ctx sdk.Context) (list []types.CmpHostResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpHostResultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CmpHostResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
