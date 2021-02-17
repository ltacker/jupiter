package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHaha{}

func NewMsgCreateHaha(creator string, text string) *MsgCreateHaha {
	return &MsgCreateHaha{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHaha) Route() string {
	return RouterKey
}

func (msg *MsgCreateHaha) Type() string {
	return "CreateHaha"
}

func (msg *MsgCreateHaha) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHaha) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHaha) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHaha{}

func NewMsgUpdateHaha(creator string, id string, text string) *MsgUpdateHaha {
	return &MsgUpdateHaha{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHaha) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHaha) Type() string {
	return "UpdateHaha"
}

func (msg *MsgUpdateHaha) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHaha) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHaha) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHaha{}

func NewMsgDeleteHaha(creator string, id string) *MsgDeleteHaha {
	return &MsgDeleteHaha{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHaha) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHaha) Type() string {
	return "DeleteHaha"
}

func (msg *MsgDeleteHaha) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHaha) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHaha) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
