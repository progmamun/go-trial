package main

import (
	"fmt"
	"math/big"

	// "math/rand"

	"crypto/rand"
)

func main() {
	fmt.Println("welcome to maths in golang")

	//random numbar
	//rand.seed(7)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(10) + 1)

	//random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(myRandomNum)
}
