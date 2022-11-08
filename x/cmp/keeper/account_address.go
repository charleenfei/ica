package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/cmp/types"
)

// SetAccountAddress set a specific accountAddress in the store from its index
func (k Keeper) SetAccountAddress(ctx sdk.Context, accountAddress types.AccountAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddressKeyPrefix))
	b := k.cdc.MustMarshal(&accountAddress)
	store.Set(types.AccountAddressKey(
		accountAddress.Index,
	), b)
}

// GetAccountAddress returns a accountAddress from its index
func (k Keeper) GetAccountAddress(
	ctx sdk.Context,
	index string,

) (val types.AccountAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddressKeyPrefix))

	b := store.Get(types.AccountAddressKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountAddress removes a accountAddress from the store
func (k Keeper) RemoveAccountAddress(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddressKeyPrefix))
	store.Delete(types.AccountAddressKey(
		index,
	))
}

// GetAllAccountAddress returns all accountAddress
func (k Keeper) GetAllAccountAddress(ctx sdk.Context) (list []types.AccountAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
