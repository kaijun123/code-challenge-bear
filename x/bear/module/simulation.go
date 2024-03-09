package bear

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"bear/testutil/sample"
	bearsimulation "bear/x/bear/simulation"
	"bear/x/bear/types"
)

// avoid unused import issue
var (
	_ = bearsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateBear = "op_weight_msg_create_bear"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBear int = 100

	opWeightMsgCreateBear = "op_weight_msg_create_bear"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBear int = 100

	opWeightMsgCreateBear = "op_weight_msg_create_bear"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBear int = 100

	opWeightMsgCreateBear = "op_weight_msg_create_bear"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBear int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bearGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bearGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBear int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBear, &weightMsgCreateBear, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBear = defaultWeightMsgCreateBear
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBear,
		bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBear int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBear, &weightMsgCreateBear, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBear = defaultWeightMsgCreateBear
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBear,
		bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBear int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBear, &weightMsgCreateBear, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBear = defaultWeightMsgCreateBear
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBear,
		bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBear int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBear, &weightMsgCreateBear, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBear = defaultWeightMsgCreateBear
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBear,
		bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBear,
			defaultWeightMsgCreateBear,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBear,
			defaultWeightMsgCreateBear,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBear,
			defaultWeightMsgCreateBear,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBear,
			defaultWeightMsgCreateBear,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bearsimulation.SimulateMsgCreateBear(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
