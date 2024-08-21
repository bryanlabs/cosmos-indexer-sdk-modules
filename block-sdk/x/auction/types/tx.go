package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgAuctionBid{}

func (msg MsgAuctionBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Bidder)}
}

func (msg MsgAuctionBid) ValidateBasic() error {
	return nil
}
