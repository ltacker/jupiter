package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHahasent{}

func NewMsgCreateHahasent(creator string, text string) *MsgCreateHahasent {
	return &MsgCreateHahasent{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHahasent) Route() string {
	return RouterKey
}

func (msg *MsgCreateHahasent) Type() string {
	return "CreateHahasent"
}

func (msg *MsgCreateHahasent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHahasent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHahasent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHahasent{}

func NewMsgUpdateHahasent(creator string, id string, text string) *MsgUpdateHahasent {
	return &MsgUpdateHahasent{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHahasent) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHahasent) Type() string {
	return "UpdateHahasent"
}

func (msg *MsgUpdateHahasent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHahasent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHahasent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHahasent{}

func NewMsgDeleteHahasent(creator string, id string) *MsgDeleteHahasent {
	return &MsgDeleteHahasent{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHahasent) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHahasent) Type() string {
	return "DeleteHahasent"
}

func (msg *MsgDeleteHahasent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHahasent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHahasent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
