package types

// DONTCOVER

import (
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "cosmossdk.io/errors"
)

// x/cmp module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrAddressDoesNotMatch = sdkerrors.Register(ModuleName, 1109, "account address must match address of message creator")
)
