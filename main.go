package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		userArg := os.Args[1]
		finalStr := "Hello, " + userArg + "!"
		fmt.Println(finalStr)
	} else {
		fmt.Println("Hello, World!")
	}

}
