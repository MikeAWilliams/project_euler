package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumIntSlice(t *testing.T) {
	require.Equal(t, 10, util.SumIntSlice([]int{5, 5}))
	require.Equal(t, 30, util.SumIntSlice([]int{10, 5, 5, 10}))
}
