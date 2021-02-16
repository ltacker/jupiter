package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendIbcmessage{}, "ibcchat/SendIbcmessage", nil)

	cdc.RegisterConcrete(&MsgCreateMessage{}, "ibcchat/CreateMessage", nil)
	cdc.RegisterConcrete(&MsgUpdateMessage{}, "ibcchat/UpdateMessage", nil)
	cdc.RegisterConcrete(&MsgDeleteMessage{}, "ibcchat/DeleteMessage", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendIbcmessage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMessage{},
		&MsgUpdateMessage{},
		&MsgDeleteMessage{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
