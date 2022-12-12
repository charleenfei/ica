package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// SetPendingSell set a specific pendingSell in the store from its index
func (k Keeper) SetPendingSell(ctx sdk.Context, pendingSell types.PendingSell) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingSellKeyPrefix))
	b := k.cdc.MustMarshal(&pendingSell)
	store.Set(types.PendingSellKey(
		pendingSell.Name,
	), b)
}

// GetPendingSell returns a pendingSell from its index
func (k Keeper) GetPendingSell(
	ctx sdk.Context,
	name string,

) (val types.PendingSell, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingSellKeyPrefix))

	b := store.Get(types.PendingSellKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePendingSell removes a pendingSell from the store
func (k Keeper) RemovePendingSell(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingSellKeyPrefix))
	store.Delete(types.PendingSellKey(
		name,
	))
}

// GetAllPendingSell returns all pendingSell
func (k Keeper) GetAllPendingSell(ctx sdk.Context) (list []types.PendingSell) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingSellKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingSell
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
