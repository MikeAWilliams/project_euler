package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDividePrecise(t *testing.T) {
	tenth := util.Divide(1, 10)
	require.Equal(t, uint64(0), tenth.Whole)
	require.Equal(t, util.Terminates, tenth.Termination)
	requireEqualSlicesU8(t, []uint8{1}, tenth.FractionDigits)
}
