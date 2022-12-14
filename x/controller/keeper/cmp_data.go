package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

// SetCmpData set a specific cmpData in the store from its index
func (k Keeper) SetCmpData(ctx sdk.Context, cmpData types.CmpData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpDataKeyPrefix))
	b := k.cdc.MustMarshal(&cmpData)
	store.Set(types.CmpDataKey(
		cmpData.Index,
	), b)
}

// GetCmpData returns a cmpData from its index
func (k Keeper) GetCmpData(
	ctx sdk.Context,
	index string,

) (val types.CmpData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpDataKeyPrefix))

	b := store.Get(types.CmpDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCmpData removes a cmpData from the store
func (k Keeper) RemoveCmpData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpDataKeyPrefix))
	store.Delete(types.CmpDataKey(
		index,
	))
}

// GetAllCmpData returns all cmpData
func (k Keeper) GetAllCmpData(ctx sdk.Context) (list []types.CmpData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CmpData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
