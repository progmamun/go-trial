package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
	// method 2
	/*
		switch tier { // switch statement
			case 1: // case statement
				fmt.Println("T-shirt")
				if age < 18 {
					break // exits the switch block
				}
				fallthrough // executes the next case
			case 2:
				fmt.Println("Mug")
				fallthrough // executes the next case
			case 3:
				fmt.Println("Sticker pack")
			default: // executed if no case is satisfied
				fmt.Println("no reward")
			}
	*/

}
