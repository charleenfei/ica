package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting a name from the store
	// whois, isFound := k.GetWhois(ctx, msg.Name)

	minPriceStruct, _ := k.GetTestmin(ctx)
	minprice, _ := strconv.Atoi(minPriceStruct.Price)
	// Set the price at which the name has to be bought if it didn't have an owner before
	// minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	// Convert price and bid strings to sdk.Coins
	// price, _ := sdk.ParseCoinsNormalized(whois.Price)
	// bid, _ := sdk.ParseCoinsNormalized(msg.Bid)
	bid, _ := strconv.Atoi(msg.Bid)

	// Convert owner and buyer address strings to sdk.AccAddress
	// owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	if bid < minprice {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
	}
	// If a name is found in store
	// if isFound {
	// If the current price is higher than the bid
	// if price.IsAllGT(bid) {
	// 	// Throw an error
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
	// }

	// Otherwise (when the bid is higher), send tokens from the buyer to the owner
	// err := k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
	// if err != nil {
	// 	return nil, err
	// }
	// } else { // If the name is not found in the store
	// If the minimum price is higher than the bid
	// if minPrice.IsAllGT(bid) {
	// 	// Throw an error
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is less than min amount")
	// }

	// Otherwise (when the bid is higher), send tokens from the buyer's account to the module's account (as a payment for the name)
	// err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
	// if err != nil {
	// 	return nil, err
	// }
	// }

	// Create an updated whois record
	// newWhois := types.Whois{
	// 	Index: msg.Name,
	// 	Name:  msg.Name,
	// 	Value: whois.Value,
	// 	Price: bid.String(),
	// 	Owner: buyer.String(),
	// }
	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: "Test ICA value",
		Price: msg.Bid,
		Owner: buyer.String(),
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CmpHostRequestEventType,
			sdk.NewAttribute(types.CmpHostRequestEventId, msg.Name),
			sdk.NewAttribute(types.CmpHostRequestEventCreator, buyer.String()),
			sdk.NewAttribute(types.CmpHostRequestItem, msg.Name),
			sdk.NewAttribute(types.CmpHostRequestBid, msg.Bid),
		),
	)
	// Write whois information to the store
	k.SetWhois(ctx, newWhois)
	return &types.MsgBuyNameResponse{}, nil
}
