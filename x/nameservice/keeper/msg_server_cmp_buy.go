package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) CmpBuy(goCtx context.Context, msg *types.MsgCmpBuy) (*types.MsgCmpBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	// Convert owner and buyer address strings to sdk.AccAddress
	// owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	uniquePendingIndex := msg.Name + ":::" + buyer.String()
	// Check if a pending buy exists in the store
	_, isFound := k.GetPendingBuy(ctx, uniquePendingIndex)
	if isFound {
		newCmpHostResult := types.CmpHostResult{
			Index:   uniquePendingIndex,
			Result:  "REJECT::" + " The same pending buy already exists" + "::::" + timestamp,
			Request: uniquePendingIndex,
		}
		// add the result to the result store
		k.SetCmpHostResult(ctx, newCmpHostResult)
		return nil, nil // dont return sdkerrors so that changes in SetCmpHostResult can be saved
		// return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "The same pending buy already exists")
	}

	buyPrice, err := strconv.Atoi(msg.Bid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Buy price is not int: "+err.Error())
	}

	pendingSell, isFound := k.GetPendingSell(ctx, msg.Name)

	if isFound {
		sellPrice, err := strconv.Atoi(pendingSell.Price)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Sell price is not int: "+err.Error())
		}

		if buyPrice != sellPrice {
			fmt.Println("Bid price does not match sell price")
			newCmpHostResult := types.CmpHostResult{
				Index:   uniquePendingIndex,
				Result:  "REJECT::" + "Bid price does not match sell price" + "::::" + timestamp,
				Request: uniquePendingIndex,
			}
			// add the result to the result store
			k.SetCmpHostResult(ctx, newCmpHostResult)
			return nil, nil // dont return sdkerrors so that changes in SetCmpHostResult can be saved
			// return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, "Bid price does not match sell price")
		}

		coins, err := sdk.ParseCoinsNormalized(pendingSell.Price + "stake")
		if err != nil {
			newCmpHostResult := types.CmpHostResult{
				Index:   uniquePendingIndex,
				Result:  "REJECT::" + "Cannot parse sell price: " + err.Error() + "::::" + timestamp,
				Request: uniquePendingIndex,
			}
			// add the result to the result store
			k.SetCmpHostResult(ctx, newCmpHostResult)
			return nil, nil // dont return sdkerrors so that changes in SetCmpHostResult can be saved
			// return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Cannot parse sell price: "+err.Error())
		}

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, coins)
		if err != nil {
			newCmpHostResult := types.CmpHostResult{
				Index:   uniquePendingIndex,
				Result:  "REJECT::" + "Insufficient Balance" + "::::" + timestamp,
				Request: uniquePendingIndex,
			}
			// add the result to the result store
			k.SetCmpHostResult(ctx, newCmpHostResult)
			return nil, nil // dont return sdkerrors so that changes in SetCmpHostResult can be saved
			// return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "Could not send coins from buyer to module")
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
