package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CmpControllerRequestAll(c context.Context, req *types.QueryAllCmpControllerRequestRequest) (*types.QueryAllCmpControllerRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cmpControllerRequests []types.CmpControllerRequest
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	cmpControllerRequestStore := prefix.NewStore(store, types.KeyPrefix(types.CmpControllerRequestKeyPrefix))

	pageRes, err := query.Paginate(cmpControllerRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var cmpControllerRequest types.CmpControllerRequest
		if err := k.cdc.Unmarshal(value, &cmpControllerRequest); err != nil {
			return err
		}
		fmt.Println("not error ")
		cmpControllerRequests = append(cmpControllerRequests, cmpControllerRequest)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryAllCmpControllerRequestResponse{CmpControllerRequest: cmpControllerRequests, Pagination: pageRes}, nil
}

func (k Keeper) CmpControllerRequest(c context.Context, req *types.QueryGetCmpControllerRequestRequest) (*types.QueryGetCmpControllerRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCmpControllerRequest(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCmpControllerRequestResponse{CmpControllerRequest: val}, nil
}
