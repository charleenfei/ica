package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmos/interchain-accounts/testutil/keeper"
	"github.com/cosmos/interchain-accounts/testutil/nullify"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCmpHostResultQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCmpHostResult(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCmpHostResultRequest
		response *types.QueryGetCmpHostResultResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCmpHostResultRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetCmpHostResultResponse{CmpHostResult: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCmpHostResultRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetCmpHostResultResponse{CmpHostResult: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCmpHostResultRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CmpHostResult(wctx, tc.request)
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

func TestCmpHostResultQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NameserviceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCmpHostResult(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCmpHostResultRequest {
		return &types.QueryAllCmpHostResultRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CmpHostResultAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CmpHostResult), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CmpHostResult),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CmpHostResultAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CmpHostResult), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CmpHostResult),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CmpHostResultAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.CmpHostResult),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CmpHostResultAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
