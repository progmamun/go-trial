package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Rename("file.txt", "../file.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}
