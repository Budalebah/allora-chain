package invariant_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
)

// log wrapper for consistent logging style
func iterationLog(t *testing.T, iteration int, a ...any) {
	t.Log(fmt.Sprint("[ITER ", iteration, "]: ", a))
}

//// warning wrapper for consistent logging style
//func iterationWarn(t *testing.T, iteration int, a ...any) {
//	iterationLog(t, iteration, "WARNING: ", a)
//}

// an actor in the simulation has a
// human readable name,
// string bech32 address,
// and an account with private key etc
type Actor struct {
	name string
	addr string
	acc  cosmosaccount.Account
}

// stringer for actor
func (a Actor) String() string {
	return a.name
}

// get the faucet name based on the seed for this test run
func getFaucetName(seed int) string {
	return "run" + strconv.Itoa(seed) + "_faucet"
}

// generates an actors name from seed and index
func getActorName(seed int, actorIndex int) string {
	return "run" + strconv.Itoa(seed) + "_actor" + strconv.Itoa(actorIndex)
}
