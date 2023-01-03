package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) QueryCmpStatus(goCtx context.Context, msg *types.MsgQueryCmpStatus) (*types.MsgQueryCmpStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	// get the cmpHostResult from store
	cmpHostResult, found := k.GetCmpHostResult(ctx, msg.Request)
	// create new response message
	var response types.MsgQueryCmpStatusResponse
	if found {
		response = types.MsgQueryCmpStatusResponse{
			Result: cmpHostResult.Result,
		}
	} else {
		response = types.MsgQueryCmpStatusResponse{
			Result: "CMP record Not Found for Request Id: " + msg.Request,
		}
	}
	return &response, nil
}
