package controller_test

import (
	"testing"

	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/controller"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CmpDataList: []types.CmpData{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CmpControllerRequestList: []types.CmpControllerRequest{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CmpControllerResultList: []types.CmpControllerResult{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ControllerKeeper(t)
	controller.InitGenesis(ctx, *k, genesisState)
	got := controller.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CmpDataList, got.CmpDataList)
	require.ElementsMatch(t, genesisState.CmpControllerRequestList, got.CmpControllerRequestList)
	require.ElementsMatch(t, genesisState.CmpControllerResultList, got.CmpControllerResultList)
	// this line is used by starport scaffolding # genesis/test/assert
}