package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	mamun := User{"Mamun", "progmamun@gmail.com", true, 21}
	fmt.Println(mamun)
	fmt.Printf("mamun details are: %+v\n", mamun)
	fmt.Printf("Name is %v and email is %v.", mamun.Name, mamun.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
