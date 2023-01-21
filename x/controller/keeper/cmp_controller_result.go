package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

// SetCmpControllerResult set a specific cmpControllerResult in the store from its index
func (k Keeper) SetCmpControllerResult(ctx sdk.Context, cmpControllerResult types.CmpControllerResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerResultKeyPrefix))
	b := k.cdc.MustMarshal(&cmpControllerResult)
	store.Set(types.CmpControllerResultKey(
		cmpControllerResult.Index,
	), b)
}

// GetCmpControllerResult returns a cmpControllerResult from its index
func (k Keeper) GetCmpControllerResult(
	ctx sdk.Context,
	index string,

) (val types.CmpControllerResult, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerResultKeyPrefix))

	b := store.Get(types.CmpControllerResultKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCmpControllerResult removes a cmpControllerResult from the store
func (k Keeper) RemoveCmpControllerResult(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerResultKeyPrefix))
	store.Delete(types.CmpControllerResultKey(
		index,
	))
}

// GetAllCmpControllerResult returns all cmpControllerResult
func (k Keeper) GetAllCmpControllerResult(ctx sdk.Context) (list []types.CmpControllerResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerResultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CmpControllerResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
