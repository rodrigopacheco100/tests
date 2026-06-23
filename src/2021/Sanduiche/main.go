package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var blockedMasks []int
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		mask := (1 << (x - 1)) | (1 << (y - 1))
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

	fmt.Println(count)
}
