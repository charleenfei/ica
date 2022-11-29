package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMinprice = "set_minprice"

var _ sdk.Msg = &MsgSetMinprice{}

func NewMsgSetMinprice(creator string, price string) *MsgSetMinprice {
	return &MsgSetMinprice{
		Creator: creator,
		Price:   price,
	}
}

func (msg *MsgSetMinprice) Route() string {
	return RouterKey
}

func (msg *MsgSetMinprice) Type() string {
	return TypeMsgSetMinprice
}

func (msg *MsgSetMinprice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMinprice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMinprice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
