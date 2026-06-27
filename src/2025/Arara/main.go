package main

import (
	"fmt"
)

func main() {

	var (
		quantityOfAraras int
		quantityOfCages  int
	)

	fmt.Scan(&quantityOfAraras, &quantityOfCages)

	if CalculateCageDistribution(quantityOfAraras, quantityOfCages) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}

func CalculateCageDistribution(quantityOfAraras, quantityOfCages int) bool {
	if quantityOfAraras == 1 {
		return true
	}

	return (quantityOfCages-1)/((quantityOfAraras-1)*5) >= 1
}
