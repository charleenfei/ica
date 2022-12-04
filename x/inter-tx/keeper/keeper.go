package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/tendermint/tendermint/libs/log"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey storetypes.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Test keeper func
func (k Keeper) Register(ctx sdk.Context, creator string, connectionId string, timeout string) {
	fmt.Println("\n")
	fmt.Println("***************************************************")
	fmt.Println("Inside intertx Register ", creator, connectionId, timeout)
	fmt.Println("***************************************************")
	fmt.Println("\n")
	k.icaControllerKeeper.RegisterInterchainAccount(ctx, connectionId, creator, "")
}

func (k Keeper) SubmitTx(ctx sdk.Context, msg *types.MsgSubmitTx) error {
	fmt.Println("\n")
	fmt.Println("***************************************************")
	fmt.Println("Inside intertx SubmitTx Keeper GoString ", msg.Owner, msg.ConnectionId, msg.Msg.GoString())
	fmt.Println("Inside intertx SubmitTx Keeper ", msg.Owner, msg.ConnectionId, msg.GetTxMsg())
	fmt.Println("***************************************************")
	fmt.Println("\n")
	portID, err := icatypes.NewControllerPortID(msg.Owner)
	if err != nil {
		return err
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{msg.GetTxMsg()})
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	_, err = k.icaControllerKeeper.SendTx(ctx, nil, msg.ConnectionId, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return err
	}
	return nil

}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
