// APACHE NOTICE: This file has been modified from the original source
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgAggregateExchangeRateVote{}
var _ sdk.Msg = &MsgAggregateExchangeRatePrevote{}
var _ sdk.Msg = &MsgDelegateFeedConsent{}

func (msg MsgAggregateExchangeRateVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Feeder)}
}

func (msg MsgAggregateExchangeRateVote) ValidateBasic() error {
	return nil
}

func (msg MsgAggregateExchangeRatePrevote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Feeder)}
}

func (msg MsgAggregateExchangeRatePrevote) ValidateBasic() error {
	return nil
}

func (msg MsgDelegateFeedConsent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Operator)}
}

func (msg MsgDelegateFeedConsent) ValidateBasic() error {
	return nil
}
