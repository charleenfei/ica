package cmp

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/interchain-accounts/testutil/sample"
	cmpsimulation "github.com/cosmos/interchain-accounts/x/cmp/simulation"
	"github.com/cosmos/interchain-accounts/x/cmp/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cmpsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAccountAddress = "op_weight_msg_account_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccountAddress int = 100

	opWeightMsgUpdateAccountAddress = "op_weight_msg_account_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccountAddress int = 100

	opWeightMsgDeleteAccountAddress = "op_weight_msg_account_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAccountAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cmpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AccountAddressList: []types.AccountAddress{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cmpGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAccountAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAccountAddress, &weightMsgCreateAccountAddress, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccountAddress = defaultWeightMsgCreateAccountAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAccountAddress,
		cmpsimulation.SimulateMsgCreateAccountAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAccountAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAccountAddress, &weightMsgUpdateAccountAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAccountAddress = defaultWeightMsgUpdateAccountAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAccountAddress,
		cmpsimulation.SimulateMsgUpdateAccountAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAccountAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAccountAddress, &weightMsgDeleteAccountAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAccountAddress = defaultWeightMsgDeleteAccountAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAccountAddress,
		cmpsimulation.SimulateMsgDeleteAccountAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
