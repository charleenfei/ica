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

	// create new response message
	var response types.MsgQueryCmpStatusResponse = types.MsgQueryCmpStatusResponse{
		Result: "OK, I got your message, from controller chain",
	}
	return &response, nil
	// return &types.MsgQueryCmpStatusResponse{}, nil
}
