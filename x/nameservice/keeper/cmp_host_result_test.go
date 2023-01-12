package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/nameservice/keeper"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCmpHostResult(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CmpHostResult {
	items := make([]types.CmpHostResult, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCmpHostResult(ctx, items[i])
	}
	return items
}

func TestCmpHostResultGet(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNCmpHostResult(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCmpHostResult(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCmpHostResultRemove(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNCmpHostResult(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCmpHostResult(ctx,
			item.Index,
		)
		_, found := keeper.GetCmpHostResult(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCmpHostResultGetAll(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNCmpHostResult(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCmpHostResult(ctx)),
	)
}
