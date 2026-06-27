package main

import (
	"fmt"
	"strings"
)

func main() {
	var fingersAmount int
	fmt.Scan(&fingersAmount)

	leftHand, rightHand := CalculateOgroFingers(fingersAmount)

	fmt.Println(leftHand, rightHand)
}

func CalculateOgroFingers(fingersAmount int) (string, string) {
	if fingersAmount == 0 {
		return "*", "*"
	} else if fingersAmount <= 5 {
		return strings.Repeat("I", fingersAmount), "*"
	} else {
		return "IIIII", strings.Repeat("I", fingersAmount-5)
	}
}
