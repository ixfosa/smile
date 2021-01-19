package main

import (
	"fmt"
	"math"
)

/*******************匿名函数********************/
/*
匿名函数是指不需要定义函数名的一种函数实现方式。1958年LISP首先采用匿名函数。

在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。

匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
 */
func main41() {
	//定义了一个名为getSqrt 的变量
	//初始化该变量时和之前的变量初始化有些不同，使用了func，func是定义函数的，
	//可是这个函数和上面说的函数最大不同就是没有函数名，也就是匿名函数。这里将一个函数当做一个变量一样的操作。
	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSqrt(4))
}

/***************************************/
//Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。
func main() {
	//// --- function variable ---
	fn := func() {
		fmt.Println("fn ", "hello world!")
	}
	fn()

	// --- function collection ---
	fns := []func(x int) {
		func(x int) {
			fmt.Println("fns ", x)
		},
		func(x int) {
			fmt.Println("fns ", x)
		},
	}
	fns[0](6)
	fns[1](9)

	// --- function as field ---
	a := struct {
		fn func(s string)
	}{
		fn: func(s string) {
			fmt.Println(s)
		},
	}
	a.fn("hello world")
}
