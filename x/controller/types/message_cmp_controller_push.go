package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCmpControllerPush = "cmp_controller_push"

var _ sdk.Msg = &MsgCmpControllerPush{}

func NewMsgCmpControllerPush(creator string, account string, kyc string, investorType string, metadata string) *MsgCmpControllerPush {
	return &MsgCmpControllerPush{
		Creator:      creator,
		Account:      account,
		Kyc:          kyc,
		InvestorType: investorType,
		Metadata:     metadata,
	}
}

func (msg *MsgCmpControllerPush) Route() string {
	return RouterKey
}

func (msg *MsgCmpControllerPush) Type() string {
	return TypeMsgCmpControllerPush
}

func (msg *MsgCmpControllerPush) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCmpControllerPush) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCmpControllerPush) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
