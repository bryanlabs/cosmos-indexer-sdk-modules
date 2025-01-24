package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgPayout{}
var _ sdk.Msg = &MsgBurnHeldBalance{}

func (msg MsgPayout) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

func (msg MsgPayout) ValidateBasic() error {
	return nil
}

func (msg MsgBurnHeldBalance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

func (msg MsgBurnHeldBalance) ValidateBasic() error {
	return nil
}
