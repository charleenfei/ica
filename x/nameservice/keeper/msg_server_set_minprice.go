package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) SetMinprice(goCtx context.Context, msg *types.MsgSetMinprice) (*types.MsgSetMinpriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	newMinprice := types.Testmin{
		Name:  "TestName",
		Price: msg.Price,
	}
	k.SetTestmin(ctx, newMinprice)
	fmt.Println("****************************************************************")
	fmt.Println("SetMinprice")
	fmt.Println(msg.String())
	return &types.MsgSetMinpriceResponse{}, nil
}
