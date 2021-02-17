package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHihisent{}

func NewMsgCreateHihisent(creator string, text string) *MsgCreateHihisent {
	return &MsgCreateHihisent{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHihisent) Route() string {
	return RouterKey
}

func (msg *MsgCreateHihisent) Type() string {
	return "CreateHihisent"
}

func (msg *MsgCreateHihisent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHihisent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHihisent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHihisent{}

func NewMsgUpdateHihisent(creator string, id string, text string) *MsgUpdateHihisent {
	return &MsgUpdateHihisent{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHihisent) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHihisent) Type() string {
	return "UpdateHihisent"
}

func (msg *MsgUpdateHihisent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHihisent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHihisent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHihisent{}

func NewMsgDeleteHihisent(creator string, id string) *MsgDeleteHihisent {
	return &MsgDeleteHihisent{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHihisent) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHihisent) Type() string {
	return "DeleteHihisent"
}

func (msg *MsgDeleteHihisent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHihisent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHihisent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
