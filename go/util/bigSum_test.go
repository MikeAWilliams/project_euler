package util_test

import (
	"maw/util"
	"testing"
)

func TestBigSum(t *testing.T) {
	requireEqualSlices(t, []int{2}, util.SumLargeInt([][]int{{1}, {1}}))
	requireEqualSlices(t, []int{1, 7}, util.SumLargeInt([][]int{{9}, {8}}))
	requireEqualSlices(t, []int{3, 7, 5}, util.SumLargeInt([][]int{{1, 9, 9}, {1, 7, 6}}))
	requireEqualSlices(t, []int{4, 2, 5}, util.SumLargeInt([][]int{{1, 9, 9}, {5, 0}, {1, 7, 6}}))
	requireEqualSlices(t, []int{1, 2, 3, 4}, util.SumLargeInt([][]int{{1, 1, 1, 1}, {1, 1, 1}, {1, 1}, {1}}))
}
