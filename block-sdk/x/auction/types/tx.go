package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgAuctionBid{}
var _ sdk.Msg = &MsgUpdateParams{}

func (msg MsgAuctionBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Bidder)}
}

func (msg MsgAuctionBid) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

func (msg MsgUpdateParams) ValidateBasic() error {
	return nil
}
