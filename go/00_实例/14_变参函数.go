package main

import "fmt"

func sum(nums ...int) int {
	fmt.Printf("%T \n", nums)
	fmt.Print(nums, "  ")
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	res1 := sum(1, 2, 3)
	fmt.Println("res1 = ", res1) // []int  [1 2 3]  res1 =  6

	res2 := sum(1,2 ,3 ,4, 5)
	fmt.Println("res2 = ", res2)


	nums := []int{1, 2, 3}
	res3 := sum(nums...)
	fmt.Println("res3 = ", res3) // [1 2 3]  res3 =  6
}
