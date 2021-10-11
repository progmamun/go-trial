package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Remove("file.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}
