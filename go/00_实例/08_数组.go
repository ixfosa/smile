package main

import "fmt"

// 在 Go 中，数组 是一个固定长度的数列。
func main() {
	// 数组默认是零值的，对于 int 数组来说也就是 0。
	var a [5]int
	fmt.Println("a = ", a) // a =  [0 0 0 0 0]

	// 使用 array[index] = value 语法来设置数组指定位置的值，或者用 array[index] 得到值
	a[1] = 1
	fmt.Println("set:", a) // set: [0 1 0 0 0]
	fmt.Println("get:", a[1]) // get: 1

	// 使用内置函数 len 返回数组的长度
	fmt.Println("len = ", len(a)) // len =  5

	// 在一行内初始化一个数组
	var b [2]string = [2]string{"fo", "zi"}
	fmt.Println("b = ", b) // b =  [fo zi]

	c := [3]int{0, 1, 2}
	fmt.Println("c = ", c) // c =  [0 1 2]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0 1 2] [1 2 3]]
}
