package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCmpBuy = "cmp_buy"

var _ sdk.Msg = &MsgCmpBuy{}

func NewMsgCmpBuy(creator string, name string, bid string, metadata string) *MsgCmpBuy {
	return &MsgCmpBuy{
		Creator:  creator,
		Name:     name,
		Bid:      bid,
		Metadata: metadata,
	}
}

func (msg *MsgCmpBuy) Route() string {
	return RouterKey
}

func (msg *MsgCmpBuy) Type() string {
	return TypeMsgCmpBuy
}

func (msg *MsgCmpBuy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCmpBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCmpBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
