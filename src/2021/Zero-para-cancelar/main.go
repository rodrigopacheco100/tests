package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	entriesLength, _ := strconv.Atoi(scanner.Text())

	entries := []int{}

	for range entriesLength {
		scanner.Scan()

		value, _ := strconv.Atoi(scanner.Text())

		if value == 0 {
			entries = entries[:len(entries)-1]
		} else {
			entries = append(entries, value)
		}
	}

	sum := 0
	for _, entry := range entries {
		sum += entry
	}

	fmt.Println(sum)
}
