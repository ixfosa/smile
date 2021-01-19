package main

import "fmt"

/*******************函数返回值********************//*

"_"标识符，用来忽略函数的某个返回值

Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。

返回值的名称应当具有一定的意义，可以作为文档使用。

没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。

Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
 */

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
func main31() {
	var a, b int = 2, 6

	c := add(a, b)

	sum, avg := calc(a, b)

	fmt.Println("c = ", c) //c =  8
	fmt.Println("sum = ", sum, "avg = ", avg) //sum =  8 avg =  4

}



/*******************多返回值可直接作为其他函数调用实参。********************/
func test() (int, int)  {
	return 1, 2
}

func sum(args ...int) int {
	var res int
	for _, v := range args {
		res += v
	}
	return res
}

func main32()  {
	fmt.Println(add(test()))
	fmt.Println(sum(test()))
}


/*****************命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。**********************/

func test4(x, y int) (z int) {
	//命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
	z = x + y
	return
}

//命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
func test5(x, y int) (z int) {
	//var z = x +y   //// 不能在一个级别，引发 "z redeclared in this block" 错误。
	{
		var z = x + y
		//return   // Error: z is shadowed during return
		return z  // 必须显式返回。
	}

}

//命名返回参数允许 defer 延迟调用通过闭包读取和修改。
func test6(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func test10(x, y int) (int) {
	z := x + y
	defer func() {
		z += 100
	}()

	return z
}

//显式 return 返回前，会先修改命名返回参数。
func test7(x, y int) (z int) {
	defer func() {
		fmt.Println("test7", z) //test7 109
	}()
	z = x + y
	return z + 100  // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
	fmt.Println("test4", test4(1, 2))
	fmt.Println("test5", test5(1, 7)) //8

	fmt.Println("test6", test6(1, 7)) //test6 108

	fmt.Println("test7", test7(1, 8)) //test7 109

	fmt.Println("test10", test10(1, 8)) ////test10 9
}


/***************************************//*


 */




/***************************************//*


 */


/***************************************//*


 */