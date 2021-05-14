package server

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/interchainberlin/metachain/x/group"
	"github.com/interchainberlin/metachain/x/group/simulation"
)

// WeightedOperations returns all the group module operations with their respective weights.
func (s serverImpl) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {

	queryClient := group.NewQueryClient(s.key)
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc,
		s.accKeeper, s.bankKeeper, queryClient,
	)
}
