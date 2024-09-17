package terraclassic

import (
	marketTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/market/types"
	oracleTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/oracle/types"
	wasmTypes "github.com/DefiantLabs/cosmos-indexer-modules/terra-classic/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var marketTypeMap = map[string]sdk.Msg{
	"/terra.market.v1beta1.MsgSwap":     &marketTypes.MsgSwap{},
	"/terra.market.v1beta1.MsgSwapSend": &marketTypes.MsgSwapSend{},
}

func GetMarketTypeMap() map[string]sdk.Msg {
	return marketTypeMap
}

var wasmTypeMap = map[string]sdk.Msg{
	"/terra.wasm.v1beta1.MsgExecuteContract":     &wasmTypes.MsgExecuteContract{},
	"/terra.wasm.v1beta1.MsgInstantiateContract": &wasmTypes.MsgInstantiateContract{},
	"/terra.wasm.v1beta1.MsgMigrateContract":     &wasmTypes.MsgMigrateContract{},
	"/terra.wasm.v1beta1.MsgMigrateCode":         &wasmTypes.MsgMigrateCode{},
	"/terra.wasm.v1beta1.MsgStoreCode":           &wasmTypes.MsgStoreCode{},
	"/terra.wasm.v1beta1.MsgUpdateContractAdmin": &wasmTypes.MsgUpdateContractAdmin{},
	"/terra.wasm.v1beta1.MsgClearContractAdmin":  &wasmTypes.MsgClearContractAdmin{},
}

func GetWasmTypeMap() map[string]sdk.Msg {
	return wasmTypeMap
}

var oracleTypeMap = map[string]sdk.Msg{
	"/terra.oracle.v1beta1.MsgAggregateExchangeRateVote":    &oracleTypes.MsgAggregateExchangeRateVote{},
	"/terra.oracle.v1beta1.MsgAggregateExchangeRatePrevote": &oracleTypes.MsgAggregateExchangeRatePrevote{},
	"/terra.oracle.v1beta1.MsgDelegateFeedConsent":          &oracleTypes.MsgDelegateFeedConsent{},
}

func GetOracleTypeMap() map[string]sdk.Msg {
	return oracleTypeMap
}

func GetTerraClassicTypeMap() map[string]sdk.Msg {

	var fullTypeMap = make(map[string]sdk.Msg)

	for k, v := range GetMarketTypeMap() {
		fullTypeMap[k] = v
	}

	for k, v := range GetWasmTypeMap() {
		fullTypeMap[k] = v
	}

	for k, v := range GetOracleTypeMap() {
		fullTypeMap[k] = v
	}

	return fullTypeMap
}
