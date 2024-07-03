package invariant_test

import (
	testCommon "github.com/allora-network/allora-chain/test/common"
)

// Use actor to create a new topic
func createTopic(m *testCommon.TestConfig, actor Actor, data *SimulationData, iteration int) error {
	iterationLog(m.T, iteration, actor, " is creating a new topic")
	return nil
}
