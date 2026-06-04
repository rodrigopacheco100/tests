package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("batata"), 10)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(hash))

	compareResult := bcrypt.CompareHashAndPassword(hash, []byte("batata"))
	isSamePassword := compareResult == nil

	fmt.Println("comparing hash with batata, result =", isSamePassword)

	compareResult = bcrypt.CompareHashAndPassword(hash, []byte("potato"))
	isSamePassword = compareResult == nil

	fmt.Println("comparing hash with potato, result =", isSamePassword)
}
