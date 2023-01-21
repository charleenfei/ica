package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CmpControllerResultAll(c context.Context, req *types.QueryAllCmpControllerResultRequest) (*types.QueryAllCmpControllerResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cmpControllerResults []types.CmpControllerResult
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cmpControllerResultStore := prefix.NewStore(store, types.KeyPrefix(types.CmpControllerResultKeyPrefix))

	pageRes, err := query.Paginate(cmpControllerResultStore, req.Pagination, func(key []byte, value []byte) error {
		var cmpControllerResult types.CmpControllerResult
		if err := k.cdc.Unmarshal(value, &cmpControllerResult); err != nil {
			return err
		}

		cmpControllerResults = append(cmpControllerResults, cmpControllerResult)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCmpControllerResultResponse{CmpControllerResult: cmpControllerResults, Pagination: pageRes}, nil
}

func (k Keeper) CmpControllerResult(c context.Context, req *types.QueryGetCmpControllerResultRequest) (*types.QueryGetCmpControllerResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCmpControllerResult(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCmpControllerResultResponse{CmpControllerResult: val}, nil
}
