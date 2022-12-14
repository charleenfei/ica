package keeper

import (
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
