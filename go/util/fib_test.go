package util_test

import (
	"maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	require.Equal(t, 1, util.Fib(1))
	require.Equal(t, 1, util.Fib(2))
	require.Equal(t, 2, util.Fib(3))
	require.Equal(t, 5, util.Fib(5))
	require.Equal(t, 144, util.Fib(12))
}

func TestFibBig(t *testing.T) {
	//require.Equal(t, int64(1), util.FibBig(1).Int64())
	//require.Equal(t, int64(1), util.FibBig(2).Int64())
	//require.Equal(t, int64(2), util.FibBig(3).Int64())
	require.Equal(t, int64(5), util.FibBig(5).Int64())
	//require.Equal(t, int64(144), util.FibBig(12).Int64())
}
