package manifestledger

import (
	manifestTypes "github.com/DefiantLabs/cosmos-indexer-modules/manifest-ledger/x/manifest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var manifestLedgerTypeMap = map[string]sdk.Msg{
	"/liftedinit.manifest.v1.MsgPayout":          &manifestTypes.MsgPayout{},
	"/liftedinit.manifest.v1.MsgBurnHeldBalance": &manifestTypes.MsgBurnHeldBalance{},
}

func GetManifestLedgerTypeMap() map[string]sdk.Msg {
	return manifestLedgerTypeMap
}
