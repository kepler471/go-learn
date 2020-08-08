package main

import (
	"fmt"
	"os"
)

// TODO move exercises to the programs they are built from
func main() {
	// prints entire output from os.Args
	fmt.Println(os.Args)
	// prints first arg
	fmt.Println(os.Args[0])
	// print each arg passed to echo, and its index
	fmt.Println("~~~~~~~~~~~")
	for n, i := range os.Args[1:] {
		fmt.Println(n, i)
	}
}
