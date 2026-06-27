package main

import (
	"testing"
)

type ogroTestCase struct {
	input         int
	expectedLeft  string
	expectedRight string
}

func TestOgro(t *testing.T) {
	testsSuite := []ogroTestCase{
		{input: 0, expectedLeft: "*", expectedRight: "*"},
		{input: 1, expectedLeft: "I", expectedRight: "*"},
		{input: 3, expectedLeft: "III", expectedRight: "*"},
		{input: 5, expectedLeft: "IIIII", expectedRight: "*"},
		{input: 6, expectedLeft: "IIIII", expectedRight: "I"},
		{input: 8, expectedLeft: "IIIII", expectedRight: "III"},
		{input: 10, expectedLeft: "IIIII", expectedRight: "IIIII"},
	}

	for _, testCase := range testsSuite {
		if leftHand, rightHand := CalculateOgroFingers(testCase.input); leftHand != testCase.expectedLeft || rightHand != testCase.expectedRight {
			t.Errorf("CalculateOgroFingers(%d) = %s %s, expected ('%s', '%s')", testCase.input, leftHand, rightHand, testCase.expectedLeft, testCase.expectedRight)
		}
	}
}
