package poa

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var POATypeMap = map[string]sdk.Msg{
	"/strangelove_ventures.poa.v1.MsgSetPower":            &MsgSetPower{},
	"/strangelove_ventures.poa.v1.MsgRemoveValidator":     &MsgRemoveValidator{},
	"/strangelove_ventures.poa.v1.MsgRemovePending":       &MsgRemovePending{},
	"/strangelove_ventures.poa.v1.MsgUpdateStakingParams": &MsgUpdateStakingParams{},
	"/strangelove_ventures.poa.v1.MsgCreateValidator":     &MsgCreateValidator{},
}

func GetPOATypeMap() map[string]sdk.Msg {
	return POATypeMap
}
