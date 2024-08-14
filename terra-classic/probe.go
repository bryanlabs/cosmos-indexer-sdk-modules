// LICENSE NOTICE: This file is not sourced from the licensed code referenced in the LICENSE file of this subfolder. Its licensed by the root LICENSE file.
package terraclassic

import (
	marketTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/market/types"
	oracleTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/oracle/types"
	wasmTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/wasm/types"
	"github.com/DefiantLabs/probe/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func IncludeMarketImplementations(client *client.ChainClient) {
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.market.v1beta1.MsgSwap", (*marketTypes.MsgSwap)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.market.v1beta1.MsgSwapSend", (*marketTypes.MsgSwapSend)(nil))
}

func IncludeWASMImplementations(client *client.ChainClient) {
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgExecuteContract", (*wasmTypes.MsgExecuteContract)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgInstantiateContract", (*wasmTypes.MsgInstantiateContract)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgMigrateContract", (*wasmTypes.MsgMigrateContract)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgMigrateCode", (*wasmTypes.MsgMigrateCode)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgStoreCode", (*wasmTypes.MsgStoreCode)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgUpdateContractAdmin", (*wasmTypes.MsgUpdateContractAdmin)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.wasm.v1beta1.MsgClearContractAdmin", (*wasmTypes.MsgClearContractAdmin)(nil))
}

func IncludeOracleImplementations(client *client.ChainClient) {
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.oracle.v1beta1.MsgAggregateExchangeRateVote", (*oracleTypes.MsgAggregateExchangeRateVote)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.oracle.v1beta1.MsgAggregateExchangeRatePrevote", (*oracleTypes.MsgAggregateExchangeRatePrevote)(nil))
	client.Codec.ProbeInterfaceRegistry.RegisterCustomTypeURL((*sdk.Msg)(nil), "/terra.oracle.v1beta1.MsgDelegateFeedConsent", (*oracleTypes.MsgDelegateFeedConsent)(nil))
}
