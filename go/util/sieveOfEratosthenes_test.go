package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSieve(t *testing.T) {
	below10 := util.GetPrimesBelowFlags(10)
	require.Equal(t, 4, util.CountPrimeFlags(below10))
	below100 := util.GetPrimesBelowFlags(100)
	require.Equal(t, 25, util.CountPrimeFlags(below100))
	below1000 := util.GetPrimesBelowFlags(1000)
	require.Equal(t, 168, util.CountPrimeFlags(below1000))
	below10000 := util.GetPrimesBelowFlags(10000)
	require.Equal(t, 1229, util.CountPrimeFlags(below10000))
	below10000000 := util.GetPrimesBelowFlags(10000000)
	require.Equal(t, 664579, util.CountPrimeFlags(below10000000))
}

func BenchmarkSieve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		util.GetPrimesBelowFlags(10000)
	}
}
