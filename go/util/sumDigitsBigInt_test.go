package util_test

import (
    "math/big"
    "maw/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumDigitsBitInt(t *testing.T) {
	require.Equal(t, int64(10), util.SumDigitsBigInt(big.NewInt(91)))
	require.Equal(t, int64(36), util.SumDigitsBigInt(big.NewInt(9999)))
}