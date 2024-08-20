// APACHE NOTICE: This file has been modified from the original source
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgExecuteContract{}
var _ sdk.Msg = &MsgInstantiateContract{}
var _ sdk.Msg = &MsgMigrateContract{}
var _ sdk.Msg = &MsgMigrateCode{}
var _ sdk.Msg = &MsgStoreCode{}
var _ sdk.Msg = &MsgUpdateContractAdmin{}
var _ sdk.Msg = &MsgClearContractAdmin{}

func (msg MsgExecuteContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgExecuteContract) ValidateBasic() error {
	return nil
}

func (msg MsgInstantiateContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgInstantiateContract) ValidateBasic() error {
	return nil
}

func (msg MsgMigrateContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Admin)}
}

func (msg MsgMigrateContract) ValidateBasic() error {
	return nil
}

func (msg MsgMigrateCode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgMigrateCode) ValidateBasic() error {
	return nil
}

func (msg MsgStoreCode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgStoreCode) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateContractAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Admin)}
}

func (msg MsgUpdateContractAdmin) ValidateBasic() error {
	return nil
}

func (msg MsgClearContractAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Admin)}
}

func (msg MsgClearContractAdmin) ValidateBasic() error {
	return nil
}
