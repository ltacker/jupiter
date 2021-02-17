package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendHihimes{}

func NewMsgSendHihimes(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
) *MsgSendHihimes {
	return &MsgSendHihimes{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

func (msg *MsgSendHihimes) Route() string {
	return RouterKey
}

func (msg *MsgSendHihimes) Type() string {
	return "SendHihimes"
}

func (msg *MsgSendHihimes) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendHihimes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendHihimes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
