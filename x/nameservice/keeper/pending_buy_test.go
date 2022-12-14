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

func createNPendingBuy(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingBuy {
	items := make([]types.PendingBuy, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPendingBuy(ctx, items[i])
	}
	return items
}

func TestPendingBuyGet(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingBuy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPendingBuy(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPendingBuyRemove(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingBuy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePendingBuy(ctx,
			item.Index,
		)
		_, found := keeper.GetPendingBuy(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPendingBuyGetAll(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	items := createNPendingBuy(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPendingBuy(ctx)),
	)
}
