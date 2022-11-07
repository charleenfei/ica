package keeper

import (
	"github.com/cosmos/interchain-accounts/x/cmp/types"
)

var _ types.QueryServer = Keeper{}
