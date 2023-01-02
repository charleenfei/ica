package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"github.com/gogo/protobuf/proto"
)

func (k msgServer) CmpControllerCallback(goCtx context.Context, msg *types.MsgCmpControllerCallback) (*types.MsgCmpControllerCallbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting CMPcontrollerRequest from the store
	pendingRequest, isFound := k.GetCmpControllerRequest(ctx, msg.Request)

	// If the request is not found, throw an error
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Request doesn't exist")
	}

	// If the message sender address is not the registered oracle address, throw error
	// Temporarily disable
	// if !(msg.Creator == Oracle_address) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized action, only authorized oracle can callback")
	// }

	// Execute logic of the CMP protocol, yes/no
	if msg.Result == "OK" || msg.Result == "YES" {
		// Unmarshall the packet data
		var packageData icatypes.InterchainAccountPacketData
		proto.Unmarshal(pendingRequest.Data, &packageData)
		fmt.Println("check existing controller request ", packageData.Type, packageData.Data)
		// Send the packet to the interchain account
		k.intertxKeeper.SubmitRawTx(ctx, pendingRequest.Owner, pendingRequest.Metadata, packageData)
	}

	// remove the pending request
	k.RemoveCmpControllerRequest(ctx, msg.Request)

	return &types.MsgCmpControllerCallbackResponse{}, nil
}
