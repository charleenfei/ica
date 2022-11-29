package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/nameservice/keeper"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func createTestTestmin(keeper *keeper.Keeper, ctx sdk.Context) types.Testmin {
	item := types.Testmin{}
	keeper.SetTestmin(ctx, item)
	return item
}

func TestTestminGet(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	item := createTestTestmin(keeper, ctx)
	rst, found := keeper.GetTestmin(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTestminRemove(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	createTestTestmin(keeper, ctx)
	keeper.RemoveTestmin(ctx)
	_, found := keeper.GetTestmin(ctx)
	require.False(t, found)
}
