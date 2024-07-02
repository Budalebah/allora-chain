package invariant_test

import (
	"fmt"
	"strconv"

	cosmossdk_io_math "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	testCommon "github.com/allora-network/allora-chain/test/common"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
)

// an actor in the simulation has a
// human readable name,
// string bech32 address,
// and an account with private key etc
type Actor struct {
	name string
	addr string
	acc  cosmosaccount.Account
}

// get the faucet name based on the seed for this test run
func getFaucetName(seed int) string {
	return "invariant" + strconv.Itoa(seed) + "_faucet"
}

// generates an actors name from seed and index
func getActorName(seed int, actorIndex int) string {
	return "invariant" + strconv.Itoa(seed) + "_actor" + strconv.Itoa(actorIndex)
}

// creates a new actor and registers them in the nodes account registry
func createNewActor(m *testCommon.TestConfig, numActors int) Actor {
	actorName := getActorName(m.Seed, numActors)
	actorAccount, _, err := m.Client.AccountRegistryCreate(actorName)
	if err != nil {
		m.T.Log("Error creating actor address: ", actorName, " - ", err)
		return Actor{}
	}
	actorAddress, err := actorAccount.Address(params.HumanCoinUnit)
	if err != nil {
		m.T.Log("Error creating actor address: ", actorName, " - ", err)
		return Actor{}
	}
	return Actor{
		name: actorName,
		addr: actorAddress,
		acc:  actorAccount,
	}
}

// creates a list of actors
func createActors(m *testCommon.TestConfig, numToCreate int) []Actor {
	actors := make([]Actor, numToCreate)
	for i := 0; i < numToCreate; i++ {
		actors[i] = createNewActor(m, i)
	}
	return actors
}

// fund every target address from the sender in amount coins
func fundActors(
	m *testCommon.TestConfig,
	sender Actor,
	targets []Actor,
	amount cosmossdk_io_math.Int,
) error {
	inputCoins := sdktypes.NewCoins(
		sdktypes.NewCoin(
			params.BaseCoinUnit,
			amount.MulRaw(int64(len(targets))),
		),
	)
	outputCoins := sdktypes.NewCoins(
		sdktypes.NewCoin(params.BaseCoinUnit, amount),
	)

	outputs := make([]banktypes.Output, len(targets))
	names := make([]string, len(targets))
	i := 0
	for _, actor := range targets {
		names[i] = actor.name
		outputs[i] = banktypes.Output{
			Address: actor.addr,
			Coins:   outputCoins,
		}
		i++
	}

	// Fund the accounts from faucet account in a single transaction
	sendMsg := &banktypes.MsgMultiSend{
		Inputs: []banktypes.Input{
			{
				Address: sender.addr,
				Coins:   inputCoins,
			},
		},
		Outputs: outputs,
	}
	_, err := m.Client.BroadcastTx(m.Ctx, sender.acc, sendMsg)
	if err != nil {
		m.T.Log("Error worker address: ", err)
		return err
	}
	m.T.Log("Funded ", len(targets), " accounts from ", sender.name, " with ", amount, " coins:", " ", names)
	return nil
}

// get the amount of money to give each actor in the simulation
// based on how much money the faucet currently has
func getPreFundAmount(m *testCommon.TestConfig, maxActors int) (cosmossdk_io_math.Int, error) {
	faucetBal, err := m.Client.QueryBank().Balance(m.Ctx, banktypes.NewQueryBalanceRequest(sdktypes.MustAccAddressFromBech32(m.FaucetAddr), params.DefaultBondDenom))
	if err != nil {
		return cosmossdk_io_math.ZeroInt(), err
	}
	// divide by 10 so you can at least run 10 runs
	amountForThisRun := faucetBal.Balance.Amount.QuoRaw(int64(10))
	ret := amountForThisRun.QuoRaw(int64(maxActors))
	if ret.Equal(cosmossdk_io_math.ZeroInt()) || ret.IsNegative() {
		return cosmossdk_io_math.ZeroInt(), fmt.Errorf("Not enough funds in faucet account to fund actors")
	}
	return ret, nil
}
