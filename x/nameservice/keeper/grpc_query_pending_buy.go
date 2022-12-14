package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingBuyAll(c context.Context, req *types.QueryAllPendingBuyRequest) (*types.QueryAllPendingBuyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pendingBuys []types.PendingBuy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pendingBuyStore := prefix.NewStore(store, types.KeyPrefix(types.PendingBuyKeyPrefix))

	pageRes, err := query.Paginate(pendingBuyStore, req.Pagination, func(key []byte, value []byte) error {
		var pendingBuy types.PendingBuy
		if err := k.cdc.Unmarshal(value, &pendingBuy); err != nil {
			return err
		}

		pendingBuys = append(pendingBuys, pendingBuy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingBuyResponse{PendingBuy: pendingBuys, Pagination: pageRes}, nil
}

func (k Keeper) PendingBuy(c context.Context, req *types.QueryGetPendingBuyRequest) (*types.QueryGetPendingBuyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPendingBuy(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPendingBuyResponse{PendingBuy: val}, nil
}
