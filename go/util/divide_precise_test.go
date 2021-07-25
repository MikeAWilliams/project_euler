package util_test

import (
	"fmt"
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireEqualDivisionResult(t *testing.T, expected, recieved util.DivisionResult) {
	failMessage := fmt.Sprintf("expected %+v, recieved %+v", expected, recieved)
	require.Equal(t, expected.Whole, recieved.Whole, failMessage)
	require.Equal(t, expected.Termination, recieved.Termination, failMessage)
	requireEqualSlicesU8(t, expected.FractionDigits, recieved.FractionDigits)
}

func TestDividePrecise(t *testing.T) {
	tenth := util.Divide(1, 10)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{1}, Termination: util.Terminates}, tenth)
}
