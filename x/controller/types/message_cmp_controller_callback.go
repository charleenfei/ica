package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCmpControllerCallback = "cmp_controller_callback"

var _ sdk.Msg = &MsgCmpControllerCallback{}

func NewMsgCmpControllerCallback(creator string, request string, result string) *MsgCmpControllerCallback {
	return &MsgCmpControllerCallback{
		Creator: creator,
		Request: request,
		Result:  result,
	}
}

func (msg *MsgCmpControllerCallback) Route() string {
	return RouterKey
}

func (msg *MsgCmpControllerCallback) Type() string {
	return TypeMsgCmpControllerCallback
}

func (msg *MsgCmpControllerCallback) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCmpControllerCallback) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCmpControllerCallback) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
