package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

func (k msgServer) SubmitTx(goCtx context.Context, msg *types.MsgSubmitTx) (*types.MsgSubmitTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var msgSubmit = intertxtypes.MsgSubmitTx{
		Owner:        msg.Creator,
		ConnectionId: msg.ConnectionId,
		Msg:          msg.Msg,
	}
	fmt.Println("\n")
	fmt.Println("-------------------------------------------")
	fmt.Println("Controller tx Msg_server SubmitTx goString ", msgSubmit.Owner, msgSubmit.ConnectionId, msgSubmit.Msg.GoString())
	fmt.Println("Controller tx Msg_server SubmitTx ", msgSubmit.Owner, msgSubmit.ConnectionId, msgSubmit.GetTxMsg())
	fmt.Println("-------------------------------------------")
	fmt.Println("\n")

	// TODO: Handling the message
	k.intertxKeeper.SubmitTx(ctx, &msgSubmit)
	return &types.MsgSubmitTxResponse{}, nil
}
