package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

func (k msgServer) CmpControllerPush(goCtx context.Context, msg *types.MsgCmpControllerPush) (*types.MsgCmpControllerPushResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// cmpEntry, isFound := k.GetCmpData(ctx, msg.Account)

	// If the message sender address is not authorized oracle, throw an error
	// Temporarily disabled
	// if !(msg.Creator == oracle) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized action, only oracle address can do this")
	// }

	newCmpEntry := types.CmpData{
		Index:        msg.Account,
		Account:      msg.Account,
		Kyc:          msg.Kyc,
		InvestorType: msg.InvestorType,
		Metadata:     msg.Metadata,
	}

	k.SetCmpData(ctx, newCmpEntry)
	return &types.MsgCmpControllerPushResponse{}, nil
}
