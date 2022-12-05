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

func CmdCmpBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cmp-buy [name] [bid] [metadata]",
		Short: "Broadcast message cmp-buy",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argBid := args[1]
			argMetadata := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCmpBuy(
				clientCtx.GetFromAddress().String(),
				argName,
				argBid,
				argMetadata,
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
