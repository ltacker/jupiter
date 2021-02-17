package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendHohomes{}, "laugh/SendHohomes", nil)

	cdc.RegisterConcrete(&MsgSendHihimes{}, "laugh/SendHihimes", nil)

	cdc.RegisterConcrete(&MsgSendHahames{}, "laugh/SendHahames", nil)

	cdc.RegisterConcrete(&MsgCreateHoho{}, "laugh/CreateHoho", nil)
	cdc.RegisterConcrete(&MsgUpdateHoho{}, "laugh/UpdateHoho", nil)
	cdc.RegisterConcrete(&MsgDeleteHoho{}, "laugh/DeleteHoho", nil)

	cdc.RegisterConcrete(&MsgCreateHihi{}, "laugh/CreateHihi", nil)
	cdc.RegisterConcrete(&MsgUpdateHihi{}, "laugh/UpdateHihi", nil)
	cdc.RegisterConcrete(&MsgDeleteHihi{}, "laugh/DeleteHihi", nil)

	cdc.RegisterConcrete(&MsgCreateHaha{}, "laugh/CreateHaha", nil)
	cdc.RegisterConcrete(&MsgUpdateHaha{}, "laugh/UpdateHaha", nil)
	cdc.RegisterConcrete(&MsgDeleteHaha{}, "laugh/DeleteHaha", nil)

	cdc.RegisterConcrete(&MsgCreateHohosent{}, "laugh/CreateHohosent", nil)
	cdc.RegisterConcrete(&MsgUpdateHohosent{}, "laugh/UpdateHohosent", nil)
	cdc.RegisterConcrete(&MsgDeleteHohosent{}, "laugh/DeleteHohosent", nil)

	cdc.RegisterConcrete(&MsgCreateHihisent{}, "laugh/CreateHihisent", nil)
	cdc.RegisterConcrete(&MsgUpdateHihisent{}, "laugh/UpdateHihisent", nil)
	cdc.RegisterConcrete(&MsgDeleteHihisent{}, "laugh/DeleteHihisent", nil)

	cdc.RegisterConcrete(&MsgCreateHahasent{}, "laugh/CreateHahasent", nil)
	cdc.RegisterConcrete(&MsgUpdateHahasent{}, "laugh/UpdateHahasent", nil)
	cdc.RegisterConcrete(&MsgDeleteHahasent{}, "laugh/DeleteHahasent", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendHohomes{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendHihimes{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendHahames{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHoho{},
		&MsgUpdateHoho{},
		&MsgDeleteHoho{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHihi{},
		&MsgUpdateHihi{},
		&MsgDeleteHihi{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHaha{},
		&MsgUpdateHaha{},
		&MsgDeleteHaha{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHohosent{},
		&MsgUpdateHohosent{},
		&MsgDeleteHohosent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHihisent{},
		&MsgUpdateHihisent{},
		&MsgDeleteHihisent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHahasent{},
		&MsgUpdateHahasent{},
		&MsgDeleteHahasent{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
