package main

import "fmt"

func main() {
	s1 := make([]int,5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println("s1： ", s1)   // s1:  [0 0 0 0 0 1 2 3]

	if s1 != nil {
		fmt.Println("s1： ", s1)   // s1:  [0 0 0 0 0 1 2 3]
	}

	var s2 []int
	if s2 == nil {
		fmt.Println("s2： nil") // s2： nil
	}

	s3 := make([]int, 0)
	s3 = append(s3, 1, 2, 3)
	fmt.Println("s3: ", s3) // s3:  [1 2 3]
}
