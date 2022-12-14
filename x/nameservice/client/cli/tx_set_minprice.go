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

func CmdSetMinprice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-minprice [price]",
		Short: "Broadcast message set-minprice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrice := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetMinprice(
				clientCtx.GetFromAddress().String(),
				argPrice,
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
