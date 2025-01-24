package osmosis

import (
	tokenFactoryTypes "github.com/DefiantLabs/cosmos-indexer-modules/osmosis/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var tokenfactoryTypeMap = map[string]sdk.Msg{
	"/osmosis.tokenfactory.v1beta1.MsgCreateDenom":      &tokenFactoryTypes.MsgCreateDenom{},
	"/osmosis.tokenfactory.v1beta1.MsgMint":             &tokenFactoryTypes.MsgMint{},
	"/osmosis.tokenfactory.v1beta1.MsgBurn":             &tokenFactoryTypes.MsgBurn{},
	"/osmosis.tokenfactory.v1beta1.MsgChangeAdmin":      &tokenFactoryTypes.MsgChangeAdmin{},
	"/osmosis.tokenfactory.v1beta1.MsgSetDenomMetadata": &tokenFactoryTypes.MsgSetDenomMetadata{},
	"/osmosis.tokenfactory.v1beta1.MsgForceTransfer":    &tokenFactoryTypes.MsgForceTransfer{},
	"/osmosis.tokenfactory.v1beta1.MsgUpdateParams":     &tokenFactoryTypes.MsgUpdateParams{},
}

func GetTokenFactoryTypeMap() map[string]sdk.Msg {
	return tokenfactoryTypeMap
}

// GetOsmosisTypeMap returns a map of osmosis message types
func GetOsmosisTypeMap() map[string]sdk.Msg {

	var fullTypeMap = make(map[string]sdk.Msg)

	for k, v := range GetTokenFactoryTypeMap() {
		fullTypeMap[k] = v
	}

	return fullTypeMap
}
