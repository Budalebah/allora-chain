package invariant_test

import (
	testCommon "github.com/allora-network/allora-chain/test/common"
)

// Every function responsible for doing a state transition
// should adhere to this function signature
type StateTransitionFunc func(m *testCommon.TestConfig, actor Actor, data *SimulationData, iteration int) error

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
// IMPORTANT: if you change getTransitionNameFromIndex function, you must
// also change this function to match it!!
func allTransitions() []StateTransitionFunc {
	return []StateTransitionFunc{
		createTopic,
	}
}

// for debugging it's helpful to be able to print the name of functions
// based on which one we picked in pickActorStateTransition
// IMPORTANT: if you change allTransitions function, you must
// also change this function to match it!!
func getTransitionNameFromIndex(index int) string {
	if index == 0 {
		return "createTopic"
	}
	return ""
}

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
func isPossibleTransition(actor Actor, data *SimulationData, transition StateTransitionFunc) bool {
	return true
}

// pickActorStateTransition picks a random state transition to take and returns which one it picked.
func pickActorStateTransition(
	m *testCommon.TestConfig,
	iteration int,
	actor Actor,
	data *SimulationData,
) StateTransitionFunc {
	transitions := allTransitions()
	for {
		randIndex := m.Client.Rand.Intn(len(transitions))
		selectedTransition := transitions[randIndex]
		if isPossibleTransition(actor, data, selectedTransition) {
			return selectedTransition
		} else {
			iterationLog(m.T, iteration, "Transition not possible: ", actor, " ", getTransitionNameFromIndex(randIndex))
		}
	}
}
