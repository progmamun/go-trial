package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	f.Close()
}
