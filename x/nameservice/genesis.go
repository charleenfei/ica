package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/nameservice/keeper"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whois
	for _, elem := range genState.WhoisList {
		k.SetWhois(ctx, elem)
	}
	// Set if defined
	if genState.Testmin != nil {
		k.SetTestmin(ctx, *genState.Testmin)
	}
	// Set all the pendingBuy
	for _, elem := range genState.PendingBuyList {
		k.SetPendingBuy(ctx, elem)
	}
	// Set all the pendingSell
	for _, elem := range genState.PendingSellList {
		k.SetPendingSell(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhoisList = k.GetAllWhois(ctx)
	// Get all testmin
	testmin, found := k.GetTestmin(ctx)
	if found {
		genesis.Testmin = &testmin
	}
	genesis.PendingBuyList = k.GetAllPendingBuy(ctx)
	genesis.PendingSellList = k.GetAllPendingSell(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
