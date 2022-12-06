package controller

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/interchain-accounts/testutil/sample"
	controllersimulation "github.com/cosmos/interchain-accounts/x/controller/simulation"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = controllersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegister = "op_weight_msg_register"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegister int = 100

	opWeightMsgRegisterIca = "op_weight_msg_register_ica"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterIca int = 100

	opWeightMsgSubmitTx = "op_weight_msg_submit_tx"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitTx int = 100

	opWeightMsgCmpControllerPush = "op_weight_msg_cmp_controller_push"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCmpControllerPush int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	controllerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&controllerGenesis)
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

	var weightMsgRegister int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegister, &weightMsgRegister, nil,
		func(_ *rand.Rand) {
			weightMsgRegister = defaultWeightMsgRegister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegister,
		controllersimulation.SimulateMsgRegister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterIca int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterIca, &weightMsgRegisterIca, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterIca = defaultWeightMsgRegisterIca
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterIca,
		controllersimulation.SimulateMsgRegisterIca(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitTx int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitTx, &weightMsgSubmitTx, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitTx = defaultWeightMsgSubmitTx
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitTx,
		controllersimulation.SimulateMsgSubmitTx(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCmpControllerPush int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCmpControllerPush, &weightMsgCmpControllerPush, nil,
		func(_ *rand.Rand) {
			weightMsgCmpControllerPush = defaultWeightMsgCmpControllerPush
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCmpControllerPush,
		controllersimulation.SimulateMsgCmpControllerPush(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
