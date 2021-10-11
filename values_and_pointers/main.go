package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Birthday() {
	u.Age++
	fmt.Println(u.Name, "turns", u.Age)
}
func main() {
	u := User{Name: "Mamun", Age: 21}
	fmt.Println(u.Name, "is now", u.Age)
	u.Birthday()
	fmt.Println(u.Name, "is now", u.Age)
}
