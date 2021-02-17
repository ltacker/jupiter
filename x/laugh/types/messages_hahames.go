package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendHahames{}

func NewMsgSendHahames(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
) *MsgSendHahames {
	return &MsgSendHahames{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

func (msg *MsgSendHahames) Route() string {
	return RouterKey
}

func (msg *MsgSendHahames) Type() string {
	return "SendHahames"
}

func (msg *MsgSendHahames) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendHahames) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendHahames) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
