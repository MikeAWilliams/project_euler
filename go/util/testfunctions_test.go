package util_test

import (
	"fmt"
	"testing"
)

func requireEqualSlicesU(t *testing.T, expected, recieved []uint) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for index, expectedValue := range expected {
		if expectedValue != recieved[index] {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}

func requireEqualSlicesU8(t *testing.T, expected, recieved []uint8) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for index, expectedValue := range expected {
		if expectedValue != recieved[index] {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}

func requireEqualSlices(t *testing.T, expected, recieved []int) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for index, expectedValue := range expected {
		if expectedValue != recieved[index] {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}

func requireEqualMaps(t *testing.T, expected, recieved map[int]int) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for eKey, eValue := range expected {
		rValue, ok := recieved[eKey]
		if !ok || rValue != eValue {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}
