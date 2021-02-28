package main

import "fmt"

func main() {
	if 5 % 2 == 1 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if 7 / 2 != 3.5 {
		fmt.Println(" 7 / 2 == 3.5")
	}

	if num := 6; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 7 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}


}
