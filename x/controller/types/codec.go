package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegister{}, "controller/Register", nil)
	cdc.RegisterConcrete(&MsgRegisterIca{}, "controller/RegisterIca", nil)
	cdc.RegisterConcrete(&MsgSubmitTx{}, "controller/SubmitTx", nil)
	cdc.RegisterConcrete(&MsgCmpControllerPush{}, "controller/CmpControllerPush", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegister{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterIca{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitTx{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCmpControllerPush{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
