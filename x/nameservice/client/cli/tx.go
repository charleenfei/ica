package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdBuyName())
	cmd.AddCommand(CmdSetName())
	cmd.AddCommand(CmdDeleteName())
	cmd.AddCommand(CmdSetMinprice())
	cmd.AddCommand(CmdCmpBuy())
	cmd.AddCommand(CmdCmpHostCallback())
	cmd.AddCommand(CmdCmpSell())
	cmd.AddCommand(CmdQueryCmpStatus())
	// this line is used by starport scaffolding # 1

	return cmd
}