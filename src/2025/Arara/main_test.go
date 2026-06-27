package main

import (
	"testing"
)

type araraTestCase struct {
	quantityOfAraras int
	quantityOfCages  int
	expected         bool
}

func TestArara(t *testing.T) {
	testsSuite := []araraTestCase{
		{quantityOfAraras: 1, quantityOfCages: 1, expected: true},
		{quantityOfAraras: 3, quantityOfCages: 11, expected: true},
		{quantityOfAraras: 4, quantityOfCages: 16, expected: true},
		{quantityOfAraras: 5, quantityOfCages: 20, expected: false},
		{quantityOfAraras: 10, quantityOfCages: 2, expected: false},
		{quantityOfAraras: 2, quantityOfCages: 4, expected: false},
	}

	for _, testCase := range testsSuite {
		if result := CalculateCageDistribution(testCase.quantityOfAraras, testCase.quantityOfCages); result != testCase.expected {
			t.Errorf("CalculateCageDistribution(%d, %d) = %v, expected %v", testCase.quantityOfAraras, testCase.quantityOfCages, result, testCase.expected)
		}
	}
}
