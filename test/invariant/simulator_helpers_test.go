package invariant_test

import (
	"fmt"
	"testing"
)

// simpler wrapper for consistent logging style
func iterationLog(t *testing.T, iteration int, a ...any) {
	t.Log(fmt.Sprintf("[ITER ", iteration, "]: ", a))
}
