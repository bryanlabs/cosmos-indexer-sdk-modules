package blocksdk

import (
	auctionTypes "github.com/DefiantLabs/cosmos-indexer-modules/block-sdk/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var AuctionTypeMap = map[string]sdk.Msg{
	"/sdk.auction.v1.MsgAuctionBid":   &auctionTypes.MsgAuctionBid{},
	"/sdk.auction.v1.MsgUpdateParams": &auctionTypes.MsgUpdateParams{},
}

func GetAuctionTypeMap() map[string]sdk.Msg {
	return AuctionTypeMap
}

func GetBlockSDKTypeMap() map[string]sdk.Msg {

	var fullTypeMap = make(map[string]sdk.Msg)

	for k, v := range GetAuctionTypeMap() {
		fullTypeMap[k] = v
	}

	return fullTypeMap
}
