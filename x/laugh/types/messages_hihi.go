package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHihi{}

func NewMsgCreateHihi(creator string, text string) *MsgCreateHihi {
	return &MsgCreateHihi{
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgCreateHihi) Route() string {
	return RouterKey
}

func (msg *MsgCreateHihi) Type() string {
	return "CreateHihi"
}

func (msg *MsgCreateHihi) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHihi) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHihi) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHihi{}

func NewMsgUpdateHihi(creator string, id string, text string) *MsgUpdateHihi {
	return &MsgUpdateHihi{
		Id:      id,
		Creator: creator,
		Text:    text,
	}
}

func (msg *MsgUpdateHihi) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHihi) Type() string {
	return "UpdateHihi"
}

func (msg *MsgUpdateHihi) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHihi) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHihi) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateHihi{}

func NewMsgDeleteHihi(creator string, id string) *MsgDeleteHihi {
	return &MsgDeleteHihi{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHihi) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHihi) Type() string {
	return "DeleteHihi"
}

func (msg *MsgDeleteHihi) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHihi) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHihi) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
