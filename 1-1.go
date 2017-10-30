package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s:", os.Args[0])
	for _, v := range os.Args[1:] {
		fmt.Printf(" %s", v)
	}
	fmt.Println()
}
