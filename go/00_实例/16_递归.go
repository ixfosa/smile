package main

import "fmt"

func main() {
	fmt.Println(fact(7))
}

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n - 1) // 5040
}