package invariant_test

import (
	testCommon "github.com/allora-network/allora-chain/test/common"
	"github.com/stretchr/testify/require"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
type SimulationData struct {
	numTopics           uint64
	maxTopics           uint64
	maxReputersPerTopic int
	maxWorkersPerTopic  int
	epochLength         int
}

// run the outer loop of the simulator
func simulate(
	m *testCommon.TestConfig,
	maxIterations int,
	numActors int,
	maxReputersPerTopic int,
	maxWorkersPerTopic int,
	topicsMax int,
	epochLength int,
) {
	// fund all actors from the faucet with some amount
	// give everybody the same amount of money to start with
	actorsList := createActors(m, numActors)
	preFundAmount, err := getPreFundAmount(m, numActors)
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
	simulationData := SimulationData{
		numTopics:           0,
		maxTopics:           uint64(topicsMax),
		maxReputersPerTopic: maxReputersPerTopic,
		maxWorkersPerTopic:  maxWorkersPerTopic,
		epochLength:         epochLength,
	}
	// every iteration, pick an actor,
	// then pick a state transition function for that actor to do
	for i := 0; i < maxIterations; i++ {
		iterationLog(m.T, i, "starting")
		require.NotNil(m.T, m.Client.Rand)
		actorNum := m.Client.Rand.Intn(numActors)
		iterationActor := actorsList[actorNum]
		stateTransitionFunc := pickActorStateTransition(m, i, iterationActor, &simulationData)
		err := stateTransitionFunc(m, iterationActor, &simulationData, i)
		if err != nil {
			m.T.Fatal(err)
		}
	}
}
