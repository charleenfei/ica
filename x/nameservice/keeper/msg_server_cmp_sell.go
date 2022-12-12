package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) CmpSell(goCtx context.Context, msg *types.MsgCmpSell) (*types.MsgCmpSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCmpSellResponse{}, nil
}
