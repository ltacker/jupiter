package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcmessage{}

func NewMsgSendIbcmessage(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	text string,
) *MsgSendIbcmessage {
	return &MsgSendIbcmessage{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Text:             text,
	}
}

func (msg *MsgSendIbcmessage) Route() string {
	return RouterKey
}

func (msg *MsgSendIbcmessage) Type() string {
	return "SendIbcmessage"
}

func (msg *MsgSendIbcmessage) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendIbcmessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendIbcmessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
