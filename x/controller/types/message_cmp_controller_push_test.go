package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/interchain-accounts/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCmpControllerPush_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCmpControllerPush
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCmpControllerPush{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCmpControllerPush{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
