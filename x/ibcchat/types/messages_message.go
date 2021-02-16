package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMessage{}

func NewMsgCreateMessage(creator string, text string) *MsgCreateMessage {
	return &MsgCreateMessage{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateMessage) Type() string {
	return "CreateMessage"
}

func (msg *MsgCreateMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMessage{}

func NewMsgUpdateMessage(creator string, id string, text string) *MsgUpdateMessage {
	return &MsgUpdateMessage{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateMessage) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMessage) Type() string {
	return "UpdateMessage"
}

func (msg *MsgUpdateMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateMessage{}

func NewMsgDeleteMessage(creator string, id string) *MsgDeleteMessage {
	return &MsgDeleteMessage{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMessage) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMessage) Type() string {
	return "DeleteMessage"
}

func (msg *MsgDeleteMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
