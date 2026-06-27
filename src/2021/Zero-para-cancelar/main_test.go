package zeroparacancelar_test

import (
	zeroparacancelar "obi-resolutions/src/2021/Zero-para-cancelar"
	"testing"
)

type zeroTestCase struct {
	values   []int
	expected int
}

func TestSumAfterCancellations(t *testing.T) {
	testsSuite := []zeroTestCase{
		{values: []int{1, 2, 0, 3}, expected: 4},
		{values: []int{1, 2, 3, 0, 0, 4}, expected: 5},
		{values: []int{1, 2, 3}, expected: 6},
		{values: []int{10}, expected: 10},
		{values: []int{5, 0, 0}, expected: 0},
		{values: []int{7, 0, 8, 0, 9}, expected: 9},
	}

	for _, testCase := range testsSuite {
		if result := zeroparacancelar.SumAfterCancellations(testCase.values); result != testCase.expected {
			t.Errorf("SumAfterCancellations(%v) = %d, expected %d", testCase.values, result, testCase.expected)
		}
	}
}
