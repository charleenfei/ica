package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/interchain-accounts/x/controller/keeper"
	"github.com/cosmos/interchain-accounts/x/controller/types"
)

func SimulateMsgCmpControllerPush(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCmpControllerPush{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CmpControllerPush simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CmpControllerPush simulation not implemented"), nil, nil
	}
}
