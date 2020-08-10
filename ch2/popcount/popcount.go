package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d -> PopCount1: %d, PopCount2: %d\n",
			i, PopCount1(uint64(i)), PopCount2(uint64(i)))
	}
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func PopCount2(x uint64) int {
	var pop byte
	for n := 0; n < 8; n++ {
		pop += pc[byte(x>>(n*8))]
	}
	return int(pop)
}
