package blocksdk

import (
	auctionTypes "github.com/DefiantLabs/cosmos-indexer-modules/block-sdk/x/auction/types"
	"github.com/DefiantLabs/probe/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var AuctionTypeMap = map[string]sdk.Msg{
	"/sdk.auction.v1.MsgAuctionBid": &auctionTypes.MsgAuctionBid{},
}

func IncludeAuctionImplementations(client *client.ChainClient) {
	for typeURL, msg := range AuctionTypeMap {
		RegisterType(client, msg, typeURL)
	}
}

func RegisterType(client *client.ChainClient, msg sdk.Msg, typeURL string) {
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), typeURL, msg)
}
