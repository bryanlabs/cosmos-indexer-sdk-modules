package terraclassic

import (
	"github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/market/types"
	"github.com/DefiantLabs/probe/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IncludeMarketImplementations(client *client.ChainClient) {
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.market.v1beta1.MsgSwap", (*types.MsgSwap)(nil))
}
