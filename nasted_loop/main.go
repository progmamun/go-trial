package main

import "fmt"

func main() {
	// for loop
	/*for { // infinite loop
	if condition {
	break // exit the loop
	}
	}
	for i < 0 { // loop with condition
	if condition {
	continue // skip current iteration and execute next
	}
	}
	for i:=0; i < 10; i++ { // loop with declaration, condition and operation
	}
	*/

	// nested loop
	var a = 9
label:
	for i := a; i < a+2; i++ {
		switch i % 3 {
		case 0:
			fmt.Println("divisible by 3")
			break label // this break the outer for loop
		default:
			fmt.Println("not divisible by 3")
		}
	}
}
