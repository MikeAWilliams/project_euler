package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func requireEqualSlices(t *testing.T, expected, recieved []int) {
	require.Equal(t, len(expected), len(recieved))
	for index, expectedValue := range expected {
		require.Equal(t, expectedValue, recieved[index])
	}
}

func TestFactors(t *testing.T) {
	requireEqualSlices(t, []int{1}, util.GetFactors(1))
	requireEqualSlices(t, []int{1, 3}, util.GetFactors(3))
	requireEqualSlices(t, []int{1, 2, 3, 6}, util.GetFactors(6))

}
