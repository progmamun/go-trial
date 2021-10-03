package main

import "fmt"

func main() {
	fmt.Println("welcome to slices lecture")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlists is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	//[1:3] [:3]
	fmt.Println(fruitList)
}
