package keeper

import (
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

var _ types.QueryServer = Keeper{}
