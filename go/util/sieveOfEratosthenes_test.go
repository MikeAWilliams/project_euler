package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPrimesBelow(t *testing.T) {
	below10 := util.GetPrimesBelow(10)
	require.Equal(t, 4, util.CountTrue(below10))
	below100 := util.GetPrimesBelow(100)
	require.Equal(t, 25, util.CountTrue(below100))
	below1000 := util.GetPrimesBelow(1000)
	require.Equal(t, 168, util.CountTrue(below1000))
	below10000 := util.GetPrimesBelow(10000)
	require.Equal(t, 1229, util.CountTrue(below10000))
}
