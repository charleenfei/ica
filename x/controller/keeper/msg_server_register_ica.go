package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

func (k msgServer) RegisterIca(goCtx context.Context, msg *types.MsgRegisterIca) (*types.MsgRegisterIcaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("\n")
	fmt.Println("-------------------------------------------")
	fmt.Println("\n")
	fmt.Println("Controller tx RegisterIca ", msg.Creator, msg.ConnectionId, msg.TimeoutHeight)
	// k.intertxKeeper.Register(ctx, "cosmos1m9l358xunhhwds0568za49mzhvuxx9uxre5tud", "connection-0", "123")
	fmt.Println("\n")
	fmt.Println("-------------------------------------------")
	fmt.Println("\n")

	k.intertxKeeper.Register(ctx, msg.Creator, msg.ConnectionId, msg.TimeoutHeight)

	return &types.MsgRegisterIcaResponse{}, nil
}
