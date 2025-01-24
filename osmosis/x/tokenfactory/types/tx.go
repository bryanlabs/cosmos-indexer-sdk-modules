package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateDenom{}
var _ sdk.Msg = &MsgMint{}
var _ sdk.Msg = &MsgBurn{}
var _ sdk.Msg = &MsgChangeAdmin{}
var _ sdk.Msg = &MsgSetDenomMetadata{}
var _ sdk.Msg = &MsgForceTransfer{}
var _ sdk.Msg = &MsgUpdateParams{}

func (msg MsgCreateDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgCreateDenom) ValidateBasic() error {
	return nil
}

func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgMint) ValidateBasic() error {
	return nil
}

func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgBurn) ValidateBasic() error {
	return nil
}

func (msg MsgChangeAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgChangeAdmin) ValidateBasic() error {
	return nil
}

func (msg MsgSetDenomMetadata) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgSetDenomMetadata) ValidateBasic() error {
	return nil
}

func (msg MsgForceTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgForceTransfer) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

func (msg MsgUpdateParams) ValidateBasic() error {
	return nil
}
