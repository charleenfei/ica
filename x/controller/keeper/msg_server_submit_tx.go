package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
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
	// Retrieve CMP data and check logic
	cmpEntry, isFound := k.GetCmpData(ctx, msg.Creator)
	if isFound {
		if strings.ToLower(cmpEntry.Kyc) == "true" || strings.ToLower(cmpEntry.Kyc) == "yes" {
			// TODO: Handling the message with extra cmp logic beyond KYC

			k.intertxKeeper.SubmitTx(ctx, &msgSubmit)

		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Crosschain tx not allowed, Account not verified, KYC required")
		}
		// Extra logic
	} else {
		data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{msgSubmit.GetTxMsg()})
		if err != nil {
			return nil, err
		}
		packetData := icatypes.InterchainAccountPacketData{
			Type: icatypes.EXECUTE_TX,
			Data: data,
		}
		packetBytes, _ := packetData.Marshal()

		newCMPControllerRequest := types.CmpControllerRequest{
			Index:    msg.Creator,
			Owner:    msg.Creator,
			Data:     packetBytes,
			Metadata: msg.ConnectionId,
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.CmpControllerRequestEventType,
				sdk.NewAttribute(types.CmpControllerRequestEventId, msg.Creator),
				sdk.NewAttribute(types.CmpControllerRequestEventCreator, msg.Creator),
				sdk.NewAttribute(types.CmpControllerRequestMetaData, "Unused metadata"),
			),
		)
		fmt.Println("CMP Controller Request index", newCMPControllerRequest.Index)
		// Write whois information to the store
		k.SetCmpControllerRequest(ctx, newCMPControllerRequest)
		return &types.MsgSubmitTxResponse{}, nil
		// return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Crosschain tx not allowed, no CMP info found")
	}

	return &types.MsgSubmitTxResponse{}, nil
}
