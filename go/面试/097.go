package main

import "fmt"

// 下面代码输出什么，请说明。
func main() {
	x := []int{1, 2, 3}
	y := [3]*int{}

	for i, v := range x {
		defer func() {
			fmt.Println(v)  // 3, 3, 3
		}()
		y[i] = &v
	}
	fmt.Println(*y[0], *y[1], *y[2]) // 3, 3, 3
}
