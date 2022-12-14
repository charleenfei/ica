package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("\n")
	fmt.Println("-------------------------------------------")
	fmt.Println("Controller tx Register ", msg.Creator, msg.Testparam)
	k.intertxKeeper.Register(ctx, msg.Creator, "connection-0", "123")
	fmt.Println("-------------------------------------------")
	fmt.Println("\n")

	return &types.MsgRegisterResponse{}, nil
}
