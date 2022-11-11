package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/cmp/keeper"
	"github.com/cosmos/interchain-accounts/x/cmp/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccountAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccountAddress {
	items := make([]types.AccountAddress, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAccountAddress(ctx, items[i])
	}
	return items
}

func TestAccountAddressGet(t *testing.T) {
	keeper, ctx := keepertest.CmpKeeper(t)
	items := createNAccountAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountAddress(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.CmpKeeper(t)
	items := createNAccountAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccountAddress(ctx,
			item.Index,
		)
		_, found := keeper.GetAccountAddress(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAccountAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.CmpKeeper(t)
	items := createNAccountAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccountAddress(ctx)),
	)
}
