package main

import "fmt"

// 下面选项正确的是？
	//- A. 1 2
	//- B. compilation error
func main() {
	if a := 1; false {
		fmt.Println(a)
	} else if b := 2; false {
		fmt.Println(b)
	} else {
		fmt.Println(a, b)  // 1, 2
	}
}
