package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCmpHostCallback() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cmp-host-callback [request] [result]",
		Short: "Broadcast message cmp-host-callback",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRequest := args[0]
			argResult := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCmpHostCallback(
				clientCtx.GetFromAddress().String(),
				argRequest,
				argResult,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
