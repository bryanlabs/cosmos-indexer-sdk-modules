package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSwap{}

func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Trader)}
}

func (msg MsgSwap) ValidateBasic() error {
	return nil
}
