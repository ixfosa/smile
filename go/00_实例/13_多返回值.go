package main

import "fmt"

func main() {
	val1, val2 := vals()
	fmt.Println("val1 = ", val1, "- val2 = ", val2) // val1 =  6 - val2 =  9

	_, val3 := vals()
	fmt.Println("val3 = ", val3) // val3 =  9
}

func vals() (int, int) {
	return 6, 9
}