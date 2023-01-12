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

func (k Keeper) CmpHostResultAll(c context.Context, req *types.QueryAllCmpHostResultRequest) (*types.QueryAllCmpHostResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cmpHostResults []types.CmpHostResult
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cmpHostResultStore := prefix.NewStore(store, types.KeyPrefix(types.CmpHostResultKeyPrefix))

	pageRes, err := query.Paginate(cmpHostResultStore, req.Pagination, func(key []byte, value []byte) error {
		var cmpHostResult types.CmpHostResult
		if err := k.cdc.Unmarshal(value, &cmpHostResult); err != nil {
			return err
		}

		cmpHostResults = append(cmpHostResults, cmpHostResult)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCmpHostResultResponse{CmpHostResult: cmpHostResults, Pagination: pageRes}, nil
}

func (k Keeper) CmpHostResult(c context.Context, req *types.QueryGetCmpHostResultRequest) (*types.QueryGetCmpHostResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCmpHostResult(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCmpHostResultResponse{CmpHostResult: val}, nil
}
