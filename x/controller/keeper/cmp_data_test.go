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

func createNCmpData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CmpData {
	items := make([]types.CmpData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCmpData(ctx, items[i])
	}
	return items
}

func TestCmpDataGet(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCmpData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCmpDataRemove(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCmpData(ctx,
			item.Index,
		)
		_, found := keeper.GetCmpData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCmpDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCmpData(ctx)),
	)
}
