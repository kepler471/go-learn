package main

import "fmt"

var p = f()

func main() {
	fmt.Println(f() == f())
	w := 1
	incr(&w)
	fmt.Println(incr(&w))
}

func f() *int {
	v := 1
	return &v
}

func incr(q *int) int {
	*q++ // increments what p points to; does not change p
	return *q
}
