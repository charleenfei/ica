package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// SetPendingBuy set a specific pendingBuy in the store from its index
func (k Keeper) SetPendingBuy(ctx sdk.Context, pendingBuy types.PendingBuy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingBuyKeyPrefix))
	b := k.cdc.MustMarshal(&pendingBuy)
	store.Set(types.PendingBuyKey(
		pendingBuy.Index,
	), b)
}

// GetPendingBuy returns a pendingBuy from its index
func (k Keeper) GetPendingBuy(
	ctx sdk.Context,
	index string,

) (val types.PendingBuy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingBuyKeyPrefix))

	b := store.Get(types.PendingBuyKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePendingBuy removes a pendingBuy from the store
func (k Keeper) RemovePendingBuy(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingBuyKeyPrefix))
	store.Delete(types.PendingBuyKey(
		index,
	))
}

// GetAllPendingBuy returns all pendingBuy
func (k Keeper) GetAllPendingBuy(ctx sdk.Context) (list []types.PendingBuy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingBuyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingBuy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
