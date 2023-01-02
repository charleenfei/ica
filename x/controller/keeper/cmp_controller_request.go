package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

// SetCmpControllerRequest set a specific cmpControllerRequest in the store from its index
func (k Keeper) SetCmpControllerRequest(ctx sdk.Context, cmpControllerRequest types.CmpControllerRequest) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerRequestKeyPrefix))
	b := k.cdc.MustMarshal(&cmpControllerRequest)
	store.Set(types.CmpControllerRequestKey(
		cmpControllerRequest.Index,
	), b)
}

// GetCmpControllerRequest returns a cmpControllerRequest from its index
func (k Keeper) GetCmpControllerRequest(
	ctx sdk.Context,
	index string,

) (val types.CmpControllerRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerRequestKeyPrefix))

	b := store.Get(types.CmpControllerRequestKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCmpControllerRequest removes a cmpControllerRequest from the store
func (k Keeper) RemoveCmpControllerRequest(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerRequestKeyPrefix))
	store.Delete(types.CmpControllerRequestKey(
		index,
	))
}

// GetAllCmpControllerRequest returns all cmpControllerRequest
func (k Keeper) GetAllCmpControllerRequest(ctx sdk.Context) (list []types.CmpControllerRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CmpControllerRequestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CmpControllerRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
