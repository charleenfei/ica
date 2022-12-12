package keeper

import (
	"context"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) CmpBuy(goCtx context.Context, msg *types.MsgCmpBuy) (*types.MsgCmpBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert owner and buyer address strings to sdk.AccAddress
	// owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	uniquePendingIndex := msg.Name + ":::" + buyer.String()
	// Check if a pending buy exists in the store
	_, isFound := k.GetPendingBuy(ctx, uniquePendingIndex)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "The same pending buy already exists")
	}

	pendingSell, isFound := k.GetPendingSell(ctx, msg.Name)

	buyPrice, err := strconv.Atoi(msg.Bid)
	if (err != nil) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Buy price is not int: " + err.Error())
	}

	if (isFound) {
		sellPrice, err := strconv.Atoi(pendingSell.Price)
		if (err != nil) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Sell price is not int: " + err.Error())
		}

		if (buyPrice < sellPrice) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, "Bid price is too low")
		}
	}

	newPendingBuy := types.PendingBuy{
		Index:    uniquePendingIndex,
		Name:     msg.Name,
		Value:    "Test ICA value",
		Price:    msg.Bid,
		Owner:    buyer.String(),
		Metadata: msg.Metadata,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CmpHostRequestEventType,
			sdk.NewAttribute(types.CmpHostRequestEventId, uniquePendingIndex),
			sdk.NewAttribute(types.CmpHostRequestEventCreator, buyer.String()),
			sdk.NewAttribute(types.CmpHostRequestItem, msg.Name),
			sdk.NewAttribute(types.CmpHostRequestBid, msg.Bid),
			sdk.NewAttribute(types.CmpHostRequestMetaData, msg.Metadata),
		),
	)
	// Write whois information to the store
	k.SetPendingBuy(ctx, newPendingBuy)

	return &types.MsgCmpBuyResponse{}, nil
}
