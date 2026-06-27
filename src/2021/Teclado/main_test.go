package teclado_test

import (
	teclado "obi-resolutions/src/2021/Teclado"
	"testing"
)

type tecladoTestCase struct {
	cellphoneNumber string
	word            string
	expected        bool
}

func TestTeclado(t *testing.T) {
	testsSuite := []tecladoTestCase{
		{cellphoneNumber: "7", word: "p", expected: true},
		{cellphoneNumber: "4663", word: "home", expected: true},
		{cellphoneNumber: "9687", word: "your", expected: true},

		{cellphoneNumber: "23992", word: "abece", expected: false},
		{cellphoneNumber: "73669", word: "phone", expected: false},
		{cellphoneNumber: "4724", word: "home", expected: false},
		{cellphoneNumber: "9677628466", word: "your", expected: false},
	}

	for _, testCase := range testsSuite {
		if result := teclado.IsMatchingCombination(testCase.cellphoneNumber, testCase.word); result != testCase.expected {
			t.Errorf("IsMatchingCombination(%s, %s) = %v, expected %v", testCase.cellphoneNumber, testCase.word, result, testCase.expected)
		}
	}
}
