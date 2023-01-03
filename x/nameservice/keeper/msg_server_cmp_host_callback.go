package keeper

import (
	"context"
	"strings"

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

	serverName := pendingBuy.Name
	var extraErrorMessages string = ""
	var err error = nil
	// Execute logic of the CMP protocol, yes/no
	if strings.HasPrefix(msg.Result, "OK") || strings.HasPrefix(msg.Result, "YES") {
		whois, ownerFound := k.GetWhois(ctx, serverName)
		if ownerFound {
			pendingSell, isSelling := k.GetPendingSell(ctx, serverName)
			if !isSelling {
				extraErrorMessages = "Name owner is not selling"
				err = sdkerrors.Wrap(sdkerrors.ErrLogic, "Name owner is not selling")
			}

			toAddr, err := sdk.AccAddressFromBech32(whois.Owner)
			if err != nil {
				extraErrorMessages = "Invalid address: " + err.Error()
				err = sdkerrors.Wrap(sdkerrors.ErrLogic, "Invalid address: "+err.Error())
			}

			coins, err := sdk.ParseCoinsNormalized(pendingSell.Price + "stake")
			if err != nil {
				extraErrorMessages = "Cannot parse coins: " + err.Error()
				err = sdkerrors.Wrap(sdkerrors.ErrLogic, "Cannot parse coins: "+err.Error())
			}

			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, coins)
			if err != nil {
				extraErrorMessages = "Could not send coins from module to seller: " + err.Error()
				err = sdkerrors.Wrap(sdkerrors.ErrLogic, "Could not send coins from module to seller: "+err.Error())
			}

			k.RemovePendingSell(ctx, serverName)
		}

		// settle buy name domain
		newWhois := types.Whois{
			Index: serverName,
			Name:  serverName,
			Value: "Test ICA value",
			Price: pendingBuy.Price,
			Owner: pendingBuy.Owner,
		}

		k.SetWhois(ctx, newWhois)
	}

	// remove the pending request
	k.RemovePendingBuy(ctx, msg.Request)
	newCmpHostResult := types.CmpHostResult{
		Index:   msg.Request,
		Result:  msg.Result + "\n" + extraErrorMessages,
		Request: msg.Request,
	}
	// add the result to the result store
	k.SetCmpHostResult(ctx, newCmpHostResult)

	return &types.MsgCmpHostCallbackResponse{}, err
}
