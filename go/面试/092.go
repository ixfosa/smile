package main

import "fmt"

const (
	q uint16 = 123
	t
	w = "abc"
	r
)
func main() {
	fmt.Printf("%T %v\n", t, t) // uint16 123
	fmt.Printf("%T %v\n", r, r) // string abc
}
