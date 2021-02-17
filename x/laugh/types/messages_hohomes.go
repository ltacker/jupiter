package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendHohomes{}

func NewMsgSendHohomes(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
) *MsgSendHohomes {
	return &MsgSendHohomes{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

func (msg *MsgSendHohomes) Route() string {
	return RouterKey
}

func (msg *MsgSendHohomes) Type() string {
	return "SendHohomes"
}

func (msg *MsgSendHohomes) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendHohomes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendHohomes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
