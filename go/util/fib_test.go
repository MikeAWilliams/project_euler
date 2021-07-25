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
