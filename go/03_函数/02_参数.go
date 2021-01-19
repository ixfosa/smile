package main

import "fmt"

/*******************函数参数********************//*

函数定义时指出，函数定义时有参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。

但当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数：

值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	func swap(x, y int) int {
		   ... ...
	}

引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

 */
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main21() {

	a, b := 1, 2
	swap(&a, &b)
	//fmt.Println("a=", a, "b=", b) //a= 2 b= 1
}


/********************可变参数*******************//*
在默认情况下，，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

注意1：	无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。
		引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。


注意2：map、slice、chan、指针、interface默认以引用的方式传递。

不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）

Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。
	func myfunc(args ...int) {    //0个或多个参数

	}

	  func add(a int, args …int) int {    //1个或多个参数

	  }

	  func add(a int, b int, args …int) int {    //2个或多个参数

	  }

注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.

任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。

用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。

	func myfunc(args ...interface{}) {

	}
 */

func test3(s string, args ...int) string {
	var x int
	for _, i := range args {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {
	fmt.Println(test3("sum: %d", 1, 2, 3)) //sum: 6

	//使用 slice 对象做变参时，必须展开。（slice...）
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(test3("sum: %d", s...)) //sum: 15  // slice... 展开slice
}



