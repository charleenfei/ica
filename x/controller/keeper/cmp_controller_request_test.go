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

func createNCmpControllerRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CmpControllerRequest {
	items := make([]types.CmpControllerRequest, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCmpControllerRequest(ctx, items[i])
	}
	return items
}

func TestCmpControllerRequestGet(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerRequest(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCmpControllerRequest(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCmpControllerRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCmpControllerRequest(ctx,
			item.Index,
		)
		_, found := keeper.GetCmpControllerRequest(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCmpControllerRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.ControllerKeeper(t)
	items := createNCmpControllerRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCmpControllerRequest(ctx)),
	)
}
