package main

import "fmt"

func main() {
	var a interface{} = "10" // 10 int print korba
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case string:
		fmt.Println("a is a string")
	}
}
