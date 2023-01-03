package types_test

import (
	"testing"

	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				WhoisList: []types.Whois{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				Testmin: &types.Testmin{
					Name:  "58",
					Price: "68",
				},
				PendingBuyList: []types.PendingBuy{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				PendingSellList: []types.PendingSell{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				CmpHostResultList: []types.CmpHostResult{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated whois",
			genState: &types.GenesisState{
				WhoisList: []types.Whois{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingBuy",
			genState: &types.GenesisState{
				PendingBuyList: []types.PendingBuy{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingSell",
			genState: &types.GenesisState{
				PendingSellList: []types.PendingSell{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated cmpHostResult",
			genState: &types.GenesisState{
				CmpHostResultList: []types.CmpHostResult{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
