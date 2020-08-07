package main

import "fmt"

func main() {
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(p != nil)
	fmt.Println(p)  // the pointer, p
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"
	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil) // "true false false"
}
