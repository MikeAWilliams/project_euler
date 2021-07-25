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
	five := util.Divide(10, 2)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 5, FractionDigits: []uint8{}, Termination: util.Terminates}, five)
	quarter := util.Divide(1, 4)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{2, 5}, Termination: util.Terminates}, quarter)

	third := util.Divide(1, 3)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{3}, Termination: util.Repeats}, third)
	sixth := util.Divide(1, 6)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{1, 6}, Termination: util.Repeats}, sixth)
	seventh := util.Divide(1, 7)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{1, 4, 2, 8, 5, 7}, Termination: util.Repeats}, seventh)
	twelfth := util.Divide(1, 12)
	requireEqualDivisionResult(t, util.DivisionResult{Whole: 0, FractionDigits: []uint8{0, 8, 3}, Termination: util.Repeats}, twelfth)

	ninEightyThree := util.Divide(1, 983)
	require.Equal(t, util.Repeats, ninEightyThree.Termination)
	require.Equal(t, 982, len(ninEightyThree.FractionDigits))
}
