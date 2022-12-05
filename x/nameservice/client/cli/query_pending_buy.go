package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
	"github.com/spf13/cobra"
)

func CmdListPendingBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pending-buy",
		Short: "list all pending-buy",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPendingBuyRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PendingBuyAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPendingBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pending-buy [index]",
		Short: "shows a pending-buy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetPendingBuyRequest{
				Index: argIndex,
			}

			res, err := queryClient.PendingBuy(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
