package sanduiche

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	blockedPairs := make([][2]int, m)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		blockedPairs[i] = [2]int{x, y}
	}

	fmt.Println(CountValidSanduiches(n, blockedPairs))
}

func CountValidSanduiches(n int, blockedPairs [][2]int) int {
	var blockedMasks []int
	for _, pair := range blockedPairs {
		mask := (1 << (pair[0] - 1)) | (1 << (pair[1] - 1))
		blockedMasks = append(blockedMasks, mask)
	}

	count := 0
	for s := 1; s < (1 << n); s++ {
		isValid := true
		for _, mask := range blockedMasks {
			if s&mask == mask {
				isValid = false
				break
			}
		}
		if isValid {
			count++
		}
	}

	return count
}
