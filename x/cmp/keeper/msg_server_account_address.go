package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/interchain-accounts/x/cmp/types"
)

func (k msgServer) CreateAccountAddress(goCtx context.Context, msg *types.MsgCreateAccountAddress) (*types.MsgCreateAccountAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	// Check if the value already exists
	_, isFound := k.GetAccountAddress(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	if msg.Creator!=msg.Index {
		return nil, sdkerrors.Wrap(types.ErrAddressDoesNotMatch, "address does not match")
	}

	var accountAddress = types.AccountAddress{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Name:      msg.Name,
		Gender:    msg.Gender,
		BirthYear: msg.BirthYear,
	}

	k.SetAccountAddress(
		ctx,
		accountAddress,
	)
	return &types.MsgCreateAccountAddressResponse{}, nil
}

func (k msgServer) UpdateAccountAddress(goCtx context.Context, msg *types.MsgUpdateAccountAddress) (*types.MsgUpdateAccountAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAccountAddress(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var accountAddress = types.AccountAddress{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Name:      msg.Name,
		Gender:    msg.Gender,
		BirthYear: msg.BirthYear,
	}

	k.SetAccountAddress(ctx, accountAddress)

	return &types.MsgUpdateAccountAddressResponse{}, nil
}

func (k msgServer) DeleteAccountAddress(goCtx context.Context, msg *types.MsgDeleteAccountAddress) (*types.MsgDeleteAccountAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAccountAddress(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAccountAddress(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteAccountAddressResponse{}, nil
}
