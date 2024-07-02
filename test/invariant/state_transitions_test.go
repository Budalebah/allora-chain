package invariant_test

import (
	testCommon "github.com/allora-network/allora-chain/test/common"
)

// Every function responsible for doing a state transition
// should adhere to this function signature
type StateTransitionFunc func(m *testCommon.TestConfig, data *SimulationData) error

// pickActorStateTransition picks a random state transition to take and returns which one it picked.
//
// The list of possible state transitions we can take are:
//
// create a new topic,
// fund a topic some more,
// stake in a reputer (delegate),
// stake as a reputer,
// register as a reputer,
// register as a worker,
// unregister as a reputer,
// unregister as a worker,
// unstake from a reputer (undelegate),
// cancel the removal of delegated stake (delegator),
// collect delegator rewards,
// unstake as a reputer,
// cancel the removal of stake (as a reputer),
// produce an inference (insert a bulk worker payload),
// produce reputation scores (insert a bulk reputer payload)
//
// state machine dependencies for valid transitions
//
// fundTopic: CreateTopic
// RegisterWorkerForTopic: CreateTopic
// RegisterReputerForTopic: CreateTopic
// stakeReputer: RegisterReputerForTopic, CreateTopic
// delegateStake: CreateTopic, RegisterReputerForTopic
// unRegisterReputer: RegisterReputerForTopic
// unRegisterWorker: RegisterWorkerForTopic
// unstakeReputer: stakeReputer
// cancelStakeRemoval: unstakeReputer
// unstakeDelegator: delegateStake
// cancelDelegateStakeRemoval: unstakeDelegator
// collectDelegatorRewards: delegateStake, fundTopic, InsertBulkWorkerPayload, InsertBulkReputerPayload
// InsertBulkWorkerPayload: RegisterWorkerForTopic, FundTopic
// InsertBulkReputerPayload: RegisterReputerForTopic, InsertBulkWorkerPayload
func pickActorStateTransition(m *testCommon.TestConfig, data *SimulationData) StateTransitionFunc {
	return createTopic
}

func createTopic(m *testCommon.TestConfig, data *SimulationData) error {
	return nil
}
