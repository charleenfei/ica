package nameservice

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/interchain-accounts/testutil/sample"
	nameservicesimulation "github.com/cosmos/interchain-accounts/x/nameservice/simulation"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nameservicesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgBuyName = "op_weight_msg_buy_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyName int = 100

	opWeightMsgSetName = "op_weight_msg_set_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetName int = 100

	opWeightMsgDeleteName = "op_weight_msg_delete_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteName int = 100

	opWeightMsgSetMinprice = "op_weight_msg_set_minprice"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetMinprice int = 100

	opWeightMsgCmpBuy = "op_weight_msg_cmp_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCmpBuy int = 100

	opWeightMsgCmpHostCallback = "op_weight_msg_cmp_host_callback"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCmpHostCallback int = 100

	opWeightMsgCmpSell = "op_weight_msg_cmp_sell"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCmpSell int = 100

	opWeightMsgQueryCmpStatus = "op_weight_msg_query_cmp_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgQueryCmpStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nameserviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nameserviceGenesis)
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

	var weightMsgBuyName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyName, &weightMsgBuyName, nil,
		func(_ *rand.Rand) {
			weightMsgBuyName = defaultWeightMsgBuyName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyName,
		nameservicesimulation.SimulateMsgBuyName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetName, &weightMsgSetName, nil,
		func(_ *rand.Rand) {
			weightMsgSetName = defaultWeightMsgSetName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetName,
		nameservicesimulation.SimulateMsgSetName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteName, &weightMsgDeleteName, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteName = defaultWeightMsgDeleteName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteName,
		nameservicesimulation.SimulateMsgDeleteName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetMinprice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetMinprice, &weightMsgSetMinprice, nil,
		func(_ *rand.Rand) {
			weightMsgSetMinprice = defaultWeightMsgSetMinprice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetMinprice,
		nameservicesimulation.SimulateMsgSetMinprice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCmpBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCmpBuy, &weightMsgCmpBuy, nil,
		func(_ *rand.Rand) {
			weightMsgCmpBuy = defaultWeightMsgCmpBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCmpBuy,
		nameservicesimulation.SimulateMsgCmpBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCmpHostCallback int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCmpHostCallback, &weightMsgCmpHostCallback, nil,
		func(_ *rand.Rand) {
			weightMsgCmpHostCallback = defaultWeightMsgCmpHostCallback
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCmpHostCallback,
		nameservicesimulation.SimulateMsgCmpHostCallback(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCmpSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCmpSell, &weightMsgCmpSell, nil,
		func(_ *rand.Rand) {
			weightMsgCmpSell = defaultWeightMsgCmpSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCmpSell,
		nameservicesimulation.SimulateMsgCmpSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgQueryCmpStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgQueryCmpStatus, &weightMsgQueryCmpStatus, nil,
		func(_ *rand.Rand) {
			weightMsgQueryCmpStatus = defaultWeightMsgQueryCmpStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgQueryCmpStatus,
		nameservicesimulation.SimulateMsgQueryCmpStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
