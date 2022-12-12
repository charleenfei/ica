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

func (k Keeper) PendingSellAll(c context.Context, req *types.QueryAllPendingSellRequest) (*types.QueryAllPendingSellResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pendingSells []types.PendingSell
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pendingSellStore := prefix.NewStore(store, types.KeyPrefix(types.PendingSellKeyPrefix))

	pageRes, err := query.Paginate(pendingSellStore, req.Pagination, func(key []byte, value []byte) error {
		var pendingSell types.PendingSell
		if err := k.cdc.Unmarshal(value, &pendingSell); err != nil {
			return err
		}

		pendingSells = append(pendingSells, pendingSell)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingSellResponse{PendingSell: pendingSells, Pagination: pageRes}, nil
}

func (k Keeper) PendingSell(c context.Context, req *types.QueryGetPendingSellRequest) (*types.QueryGetPendingSellResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPendingSell(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPendingSellResponse{PendingSell: val}, nil
}
