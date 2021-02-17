package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHoho{}

func NewMsgCreateHoho(creator string, text string) *MsgCreateHoho {
	return &MsgCreateHoho{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHoho) Route() string {
	return RouterKey
}

func (msg *MsgCreateHoho) Type() string {
	return "CreateHoho"
}

func (msg *MsgCreateHoho) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHoho) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHoho) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHoho{}

func NewMsgUpdateHoho(creator string, id string, text string) *MsgUpdateHoho {
	return &MsgUpdateHoho{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHoho) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHoho) Type() string {
	return "UpdateHoho"
}

func (msg *MsgUpdateHoho) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHoho) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHoho) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHoho{}

func NewMsgDeleteHoho(creator string, id string) *MsgDeleteHoho {
	return &MsgDeleteHoho{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHoho) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHoho) Type() string {
	return "DeleteHoho"
}

func (msg *MsgDeleteHoho) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHoho) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHoho) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
