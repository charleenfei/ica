package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgQueryCmpStatus = "query_cmp_status"

var _ sdk.Msg = &MsgQueryCmpStatus{}

func NewMsgQueryCmpStatus(creator string, request string) *MsgQueryCmpStatus {
	return &MsgQueryCmpStatus{
		Creator: creator,
		Request: request,
	}
}

func (msg *MsgQueryCmpStatus) Route() string {
	return RouterKey
}

func (msg *MsgQueryCmpStatus) Type() string {
	return TypeMsgQueryCmpStatus
}

func (msg *MsgQueryCmpStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgQueryCmpStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgQueryCmpStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
