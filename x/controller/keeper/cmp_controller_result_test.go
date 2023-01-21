package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/controller/keeper"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCmpControllerResult(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CmpControllerResult {
	items := make([]types.CmpControllerResult, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCmpControllerResult(ctx, items[i])
	}
	return items
}

func TestCmpControllerResultGet(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerResult(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCmpControllerResult(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCmpControllerResultRemove(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerResult(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCmpControllerResult(ctx,
			item.Index,
		)
		_, found := keeper.GetCmpControllerResult(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCmpControllerResultGetAll(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerResult(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCmpControllerResult(ctx)),
	)
}
