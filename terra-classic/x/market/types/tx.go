package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSwap{}
var _ sdk.Msg = &MsgSwapSend{}

func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Trader)}
}

func (msg MsgSwap) ValidateBasic() error {
	return nil
}

func (msg MsgSwapSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.FromAddress)}
}

func (msg MsgSwapSend) ValidateBasic() error {
	return nil
}
