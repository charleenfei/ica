package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/interchain-accounts/x/nameservice/keeper"
	"github.com/cosmos/interchain-accounts/x/nameservice/types"
)

func SimulateMsgCmpSell(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCmpSell{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CmpSell simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CmpSell simulation not implemented"), nil, nil
	}
}
