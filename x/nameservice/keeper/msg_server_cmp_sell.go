package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) CmpSell(goCtx context.Context, msg *types.MsgCmpSell) (*types.MsgCmpSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, isFound := k.GetWhois(ctx, msg.Name)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name is not yet registered")
	}

	if msg.Creator != whois.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are not owner of this name")
	}

	newPendingSell := types.PendingSell{
		Name:  msg.Name,
		Price: msg.Price,
	}

	k.SetPendingSell(ctx, newPendingSell)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCmpSellResponse{}, nil
}
