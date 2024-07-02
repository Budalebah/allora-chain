package invariant_test

import (
	testCommon "github.com/allora-network/allora-chain/test/common"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
type SimulationData struct {
	numTopics uint64
}

// run the outer loop of the simulator
func simulate(
	m *testCommon.TestConfig,
	maxIterations int,
	maxActors int,
	maxReputersPerTopic int,
	maxWorkersPerTopic int,
	topicsMax int,
	epochLength int,
) {
	// fund all actors from the faucet with some amount
	// give everybody the same amount of money to start with
	actorsList := createActors(m, maxActors)
	preFundAmount, err := getPreFundAmount(m, maxActors)
	if err != nil {
		m.T.Fatal(err)
	}
	faucet := Actor{
		name: getFaucetName(m.Seed),
		addr: m.FaucetAddr,
		acc:  m.FaucetAcc,
	}
	err = fundActors(
		m,
		faucet,
		actorsList,
		preFundAmount,
	)
	if err != nil {
		m.T.Fatal(err)
	}
	// every iteration, pick an actor,
	// then pick a state transition function for that actor to do
	for i := 0; i < maxIterations; i++ {
	}
}
