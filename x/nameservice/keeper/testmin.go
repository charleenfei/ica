package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// SetTestmin set testmin in the store
func (k Keeper) SetTestmin(ctx sdk.Context, testmin types.Testmin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestminKey))
	b := k.cdc.MustMarshal(&testmin)
	store.Set([]byte{0}, b)
}

// GetTestmin returns testmin
func (k Keeper) GetTestmin(ctx sdk.Context) (val types.Testmin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestminKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTestmin removes testmin from the store
func (k Keeper) RemoveTestmin(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestminKey))
	store.Delete([]byte{0})
}
