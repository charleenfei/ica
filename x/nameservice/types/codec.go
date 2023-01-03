package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(&MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(&MsgDeleteName{}, "nameservice/DeleteName", nil)
	cdc.RegisterConcrete(&MsgSetMinprice{}, "nameservice/SetMinprice", nil)
	cdc.RegisterConcrete(&MsgCmpBuy{}, "nameservice/CmpBuy", nil)
	cdc.RegisterConcrete(&MsgCmpHostCallback{}, "nameservice/CmpHostCallback", nil)
	cdc.RegisterConcrete(&MsgCmpSell{}, "nameservice/CmpSell", nil)
	cdc.RegisterConcrete(&MsgQueryCmpStatus{}, "nameservice/QueryCmpStatus", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetMinprice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCmpBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCmpHostCallback{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCmpSell{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgQueryCmpStatus{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
