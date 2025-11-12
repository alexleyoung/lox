package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("Arguments without program name:", os.Args[1:])
		os.Exit(64)
	}
}
