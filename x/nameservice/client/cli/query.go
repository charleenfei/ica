package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group nameservice queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListWhois())
	cmd.AddCommand(CmdShowWhois())
	cmd.AddCommand(CmdShowTestmin())
	cmd.AddCommand(CmdListPendingBuy())
	cmd.AddCommand(CmdShowPendingBuy())
	cmd.AddCommand(CmdListPendingSell())
	cmd.AddCommand(CmdShowPendingSell())
	cmd.AddCommand(CmdListCmpHostResult())
	cmd.AddCommand(CmdShowCmpHostResult())
	// this line is used by starport scaffolding # 1

	return cmd
}