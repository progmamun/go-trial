package main

import (
	"fmt"
	"math/rand"

	// "crypto/rand"
	"time"
)

func main() {
	fmt.Println("welcome to maths in golang")

	//random numbar
	//rand.seed(7)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10) + 1)
}
