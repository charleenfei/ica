package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCmpHostCallback = "cmp_host_callback"

var _ sdk.Msg = &MsgCmpHostCallback{}

func NewMsgCmpHostCallback(creator string, request string, result string) *MsgCmpHostCallback {
	return &MsgCmpHostCallback{
		Creator: creator,
		Request: request,
		Result:  result,
	}
}

func (msg *MsgCmpHostCallback) Route() string {
	return RouterKey
}

func (msg *MsgCmpHostCallback) Type() string {
	return TypeMsgCmpHostCallback
}

func (msg *MsgCmpHostCallback) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCmpHostCallback) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCmpHostCallback) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
