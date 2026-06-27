package main

import (
	"fmt"
	"slices"
)

func main() {
	var cellphoneNumber string
	var entriesAmount int

	fmt.Scan(&cellphoneNumber, &entriesAmount)

	var combinations = 0
	entries := make([]string, entriesAmount)

	for i := range entries {
		fmt.Scan(&entries[i])
	}

	for _, entry := range entries {
		if isValid := IsMatchingCombination(cellphoneNumber, entry); isValid {
			combinations++
		}
	}

	fmt.Println(combinations)
}

func IsMatchingCombination(
	cellphoneNumber string,
	word string,
) bool {
	mappedKeyboard := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if len(cellphoneNumber) != len(word) {
		return false
	}

	var isValid = true
	for i := range cellphoneNumber {
		cellphoneNumberDigit := cellphoneNumber[i]
		combinationDigit := word[i]

		digitForNumber := mappedKeyboard[string(cellphoneNumberDigit)]
		if !slices.Contains(digitForNumber, string(combinationDigit)) {
			isValid = false
			break
		}
	}

	return isValid
}
