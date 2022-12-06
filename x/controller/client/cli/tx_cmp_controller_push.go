package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/interchain-accounts/x/controller/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCmpControllerPush() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cmp-controller-push [account] [kyc] [investor-type] [metadata]",
		Short: "Broadcast message cmp-controller-push",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccount := args[0]
			argKyc := args[1]
			argInvestorType := args[2]
			argMetadata := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCmpControllerPush(
				clientCtx.GetFromAddress().String(),
				argAccount,
				argKyc,
				argInvestorType,
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
