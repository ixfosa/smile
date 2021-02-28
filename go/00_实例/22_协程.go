package main

import "fmt"

func fn(from string)  {
	for i := 1; i <= 5; i++ {
		fmt.Println(from, " ", i)
	}
}
func main() {
	fn("python")

	go fn("java")

	go func(msg string) {
		for i := 1; i <= 5; i++ {
			fmt.Println(msg, " ", i)
		}
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}


