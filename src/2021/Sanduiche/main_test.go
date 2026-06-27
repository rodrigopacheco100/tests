package main

import (
	"testing"
)

type sanduicheTestCase struct {
	n            int
	blockedPairs [][2]int
	expected     int
}

func TestCountValidSandwiches(t *testing.T) {
	testsSuite := []sanduicheTestCase{
		{n: 3, blockedPairs: [][2]int{}, expected: 7},
		{n: 3, blockedPairs: [][2]int{{1, 2}}, expected: 5},
		{n: 3, blockedPairs: [][2]int{{1, 2}, {2, 3}}, expected: 4},
		{n: 4, blockedPairs: [][2]int{}, expected: 15},
		{n: 1, blockedPairs: [][2]int{}, expected: 1},
		{n: 2, blockedPairs: [][2]int{{1, 2}}, expected: 2},
	}

	for _, testCase := range testsSuite {
		if result := CountValidSandwiches(testCase.n, testCase.blockedPairs); result != testCase.expected {
			t.Errorf("CountValidSandwiches(%d, %v) = %d, expected %d", testCase.n, testCase.blockedPairs, result, testCase.expected)
		}
	}
}
