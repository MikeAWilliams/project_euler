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

func TestGetPrimes(t *testing.T) {
	twentyFlags := util.GetPrimesBelowFlags(20)
	requireEqualSlicesU(t, []uint{2, 3, 5, 7, 11, 13, 17, 19}, util.GetPrimes(twentyFlags))
}

func BenchmarkSieveFlags(b *testing.B) {
	var result []bool
	for n := 0; n < b.N; n++ {
		result = util.GetPrimesBelowFlags(10000)
		util.GetPrimes(result)
	}
	result[0] = true
}

func BenchmarkSievePrimes(b *testing.B) {
	var result []uint
	flags := util.GetPrimesBelowFlags(10000)
	for n := 0; n < b.N; n++ {
		result = util.GetPrimes(flags)
	}
	result[0] = 1
}
