package poa

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSetPower{}
var _ sdk.Msg = &MsgRemoveValidator{}
var _ sdk.Msg = &MsgRemovePending{}
var _ sdk.Msg = &MsgUpdateStakingParams{}
var _ sdk.Msg = &MsgCreateValidator{}

func (msg MsgSetPower) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgSetPower) ValidateBasic() error {
	return nil
}

func (msg MsgRemoveValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgRemoveValidator) ValidateBasic() error {
	return nil
}

func (msg MsgRemovePending) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgRemovePending) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateStakingParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgUpdateStakingParams) ValidateBasic() error {
	return nil
}

func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddress)}
}

func (msg MsgCreateValidator) ValidateBasic() error {
	return nil
}
