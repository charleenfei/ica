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

func (k Keeper) CmpDataAll(c context.Context, req *types.QueryAllCmpDataRequest) (*types.QueryAllCmpDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cmpDatas []types.CmpData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cmpDataStore := prefix.NewStore(store, types.KeyPrefix(types.CmpDataKeyPrefix))

	pageRes, err := query.Paginate(cmpDataStore, req.Pagination, func(key []byte, value []byte) error {
		var cmpData types.CmpData
		if err := k.cdc.Unmarshal(value, &cmpData); err != nil {
			return err
		}

		cmpDatas = append(cmpDatas, cmpData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCmpDataResponse{CmpData: cmpDatas, Pagination: pageRes}, nil
}

func (k Keeper) CmpData(c context.Context, req *types.QueryGetCmpDataRequest) (*types.QueryGetCmpDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCmpData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCmpDataResponse{CmpData: val}, nil
}
