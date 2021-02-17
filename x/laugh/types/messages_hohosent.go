package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHohosent{}

func NewMsgCreateHohosent(creator string, text string) *MsgCreateHohosent {
	return &MsgCreateHohosent{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHohosent) Route() string {
	return RouterKey
}

func (msg *MsgCreateHohosent) Type() string {
	return "CreateHohosent"
}

func (msg *MsgCreateHohosent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHohosent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHohosent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHohosent{}

func NewMsgUpdateHohosent(creator string, id string, text string) *MsgUpdateHohosent {
	return &MsgUpdateHohosent{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHohosent) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHohosent) Type() string {
	return "UpdateHohosent"
}

func (msg *MsgUpdateHohosent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHohosent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHohosent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHohosent{}

func NewMsgDeleteHohosent(creator string, id string) *MsgDeleteHohosent {
	return &MsgDeleteHohosent{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHohosent) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHohosent) Type() string {
	return "DeleteHohosent"
}

func (msg *MsgDeleteHohosent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHohosent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHohosent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
