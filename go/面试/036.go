package main

import "fmt"

// 下面代码中，x 已声明，y 没有声明，判断每条语句的对错。
	//- 1. x, _ := f()  F
	//- 2. x, _ = f()   T
	//- 3. x, y := f()  T
	//- 4. x, y = f()   F

func f() (int, int) {
	return 1, 2
}

func main() {
	x := 0   // 声明 x

	// x, _ := f() // no new variables on left side of :=

	// x, _ = f() // 1

	x, y := f() // 1, 2


	// x, y = f() // undefined: y



	fmt.Println(x)
	fmt.Println(y)

}
