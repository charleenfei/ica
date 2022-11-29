package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func TestTestminQuery(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTestmin(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTestminRequest
		response *types.QueryGetTestminResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTestminRequest{},
			response: &types.QueryGetTestminResponse{Testmin: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Testmin(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
