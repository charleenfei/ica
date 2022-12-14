package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) CmpHostCallback(goCtx context.Context, msg *types.MsgCmpHostCallback) (*types.MsgCmpHostCallbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// Try getting name information from the store
	pendingBuy, isFound := k.GetPendingBuy(ctx, msg.Request)

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
		// settle buy name domain
		newWhois := types.Whois{
			Index: pendingBuy.Name,
			Name:  pendingBuy.Name,
			Value: "Test ICA value",
			Price: pendingBuy.Price,
			Owner: pendingBuy.Owner,
		}
		k.SetWhois(ctx, newWhois)
	}

	// remove the pending request
	k.RemovePendingBuy(ctx, msg.Request)

	return &types.MsgCmpHostCallbackResponse{}, nil
}
