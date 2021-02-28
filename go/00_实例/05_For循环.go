package main

import "fmt"

func main() {
	// 带单个循环条件
	i := 0
	for i < 3 {
		fmt.Println("i = ", i)
		i += 1
	}

	// 初始化/条件/后续形式 for 循环
	for i := 0; i <= 3; i++ {
		fmt.Println("i = ", i)
	}

	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环。
	for {
		fmt.Printf("loop")
		break
	}

}
