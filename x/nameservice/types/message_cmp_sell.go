package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCmpSell = "cmp_sell"

var _ sdk.Msg = &MsgCmpSell{}

func NewMsgCmpSell(creator string, name string, price string, metadata string) *MsgCmpSell {
	return &MsgCmpSell{
		Creator:  creator,
		Name:     name,
		Price:    price,
		Metadata: metadata,
	}
}

func (msg *MsgCmpSell) Route() string {
	return RouterKey
}

func (msg *MsgCmpSell) Type() string {
	return TypeMsgCmpSell
}

func (msg *MsgCmpSell) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCmpSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCmpSell) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
