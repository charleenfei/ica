package controller

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/controller/keeper"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the cmpData
	for _, elem := range genState.CmpDataList {
		k.SetCmpData(ctx, elem)
	}
	// Set all the cmpControllerRequest
	for _, elem := range genState.CmpControllerRequestList {
		k.SetCmpControllerRequest(ctx, elem)
	}
	// Set all the cmpControllerResult
	for _, elem := range genState.CmpControllerResultList {
		k.SetCmpControllerResult(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CmpDataList = k.GetAllCmpData(ctx)
	genesis.CmpControllerRequestList = k.GetAllCmpControllerRequest(ctx)
	genesis.CmpControllerResultList = k.GetAllCmpControllerResult(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
