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

	values := make([]int, entriesLength)
	for i := range entriesLength {
		scanner.Scan()
		values[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println(SumAfterCancellations(values))
}

func SumAfterCancellations(values []int) int {
	entries := []int{}

	for _, v := range values {
		if v == 0 {
			if len(entries) > 0 {
				entries = entries[:len(entries)-1]
			}
		} else {
			entries = append(entries, v)
		}
	}

	sum := 0
	for _, e := range entries {
		sum += e
	}

	return sum
}
