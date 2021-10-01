package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers class")

	// var ptr *int
	//*string

	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	//ref. mean &

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)
}
