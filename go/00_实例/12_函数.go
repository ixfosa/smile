package main

import "fmt"

// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
func main() {
	res := plus(1, 2)
	fmt.Println("res = ", res) // res =  3
}

func plus(a, b int) int {
	return a + b
}
