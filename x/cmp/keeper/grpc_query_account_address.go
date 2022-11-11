package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/interchain-accounts/x/cmp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountAddressAll(c context.Context, req *types.QueryAllAccountAddressRequest) (*types.QueryAllAccountAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountAddresss []types.AccountAddress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountAddressStore := prefix.NewStore(store, types.KeyPrefix(types.AccountAddressKeyPrefix))

	pageRes, err := query.Paginate(accountAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var accountAddress types.AccountAddress
		if err := k.cdc.Unmarshal(value, &accountAddress); err != nil {
			return err
		}

		accountAddresss = append(accountAddresss, accountAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

/*	accountAddresss = append(accountAddresss,
		types.AccountAddress {
			Creator:   "cosmos_address",
			Index:     "cosmos_address",
			Name:      "Alice",
			Gender:    "Female",
			BirthYear: 2001,
		},
	) */

	return &types.QueryAllAccountAddressResponse{AccountAddress: accountAddresss, Pagination: pageRes}, nil
}

func (k Keeper) AccountAddress(c context.Context, req *types.QueryGetAccountAddressRequest) (*types.QueryGetAccountAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccountAddress(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountAddressResponse{AccountAddress: val}, nil
}
