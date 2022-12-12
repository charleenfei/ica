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

func createNPendingSell(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingSell {
	items := make([]types.PendingSell, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetPendingSell(ctx, items[i])
	}
	return items
}

func TestPendingSellGet(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingSell(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPendingSell(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPendingSellRemove(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingSell(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePendingSell(ctx,
			item.Name,
		)
		_, found := keeper.GetPendingSell(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestPendingSellGetAll(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingSell(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPendingSell(ctx)),
	)
}
