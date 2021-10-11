package main

import (
	"fmt"
	"os"
)

func main() {
	// let's keep thing under 4KB
	if err := os.Truncate("file.txt", 4096); err != nil {
		fmt.Println("Error:", err)
	}
}
