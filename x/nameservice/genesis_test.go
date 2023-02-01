package nameservice_test

import (
	"testing"

	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/nameservice"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhoisList: []types.Whois{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		Testmin: &types.Testmin{
			Name:  "85",
			Price: "28",
		},
		PendingBuyList: []types.PendingBuy{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PendingSellList: []types.PendingSell{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		CmpHostResultList: []types.CmpHostResult{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, genesisState)
	got := nameservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhoisList, got.WhoisList)
	require.Equal(t, genesisState.Testmin, got.Testmin)
	require.ElementsMatch(t, genesisState.PendingBuyList, got.PendingBuyList)
	require.ElementsMatch(t, genesisState.PendingSellList, got.PendingSellList)
	require.ElementsMatch(t, genesisState.CmpHostResultList, got.CmpHostResultList)
	// this line is used by starport scaffolding # genesis/test/assert
}