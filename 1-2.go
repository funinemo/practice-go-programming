package main

import (
	"fmt"
	"os"
)

func main() {
	//	fmt.Printf("%s:", os.Args[0])
	for k, v := range os.Args[1:] {
		fmt.Printf("%d: %s\n", k, v)
	}
	//	fmt.Println()
}
