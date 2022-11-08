package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAccountAddress = "create_account_address"
	TypeMsgUpdateAccountAddress = "update_account_address"
	TypeMsgDeleteAccountAddress = "delete_account_address"
)

var _ sdk.Msg = &MsgCreateAccountAddress{}

func NewMsgCreateAccountAddress(
	creator string,
	index string,
	name string,
	gender string,
	birthYear uint64,

) *MsgCreateAccountAddress {
	return &MsgCreateAccountAddress{
		Creator:   creator,
		Index:     index,
		Name:      name,
		Gender:    gender,
		BirthYear: birthYear,
	}
}

func (msg *MsgCreateAccountAddress) Route() string {
	return RouterKey
}

func (msg *MsgCreateAccountAddress) Type() string {
	return TypeMsgCreateAccountAddress
}

func (msg *MsgCreateAccountAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAccountAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAccountAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAccountAddress{}

func NewMsgUpdateAccountAddress(
	creator string,
	index string,
	name string,
	gender string,
	birthYear uint64,

) *MsgUpdateAccountAddress {
	return &MsgUpdateAccountAddress{
		Creator:   creator,
		Index:     index,
		Name:      name,
		Gender:    gender,
		BirthYear: birthYear,
	}
}

func (msg *MsgUpdateAccountAddress) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAccountAddress) Type() string {
	return TypeMsgUpdateAccountAddress
}

func (msg *MsgUpdateAccountAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAccountAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAccountAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAccountAddress{}

func NewMsgDeleteAccountAddress(
	creator string,
	index string,

) *MsgDeleteAccountAddress {
	return &MsgDeleteAccountAddress{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteAccountAddress) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAccountAddress) Type() string {
	return TypeMsgDeleteAccountAddress
}

func (msg *MsgDeleteAccountAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAccountAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAccountAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
