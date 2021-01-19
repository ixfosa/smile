package main

import "fmt"

/******************** if 语句 由一个布尔表达式后紧跟一个或多个语句组成。*******************/
/*
 	• 可省略条件表达式括号。
    • 持初始化语句，可定义代码块局部变量。
    • 代码块左 括号必须在条件表达式尾部。

    if 布尔表达式 {
   		//在布尔表达式为 true 时执行
	}
 */
func main1() {

	var x int = 9 //定义局部变量

	//使用 if 语句判断布尔表达式
	if x < 10 {
		fmt.Printf("a 小于 10\n" )
	}
	fmt.Printf("a 的值为 : %d\n", x)
	//a 小于 10
	//a 的值为 : 9
}

func main11()  {
	x := 0
	if str := "ixfosa"; x < 0 {
		fmt.Println(str[0])
	} else if x > 0 {
		fmt.Println(str[1])
	} else {
		fmt.Println(string(str[2]))
	}
	// *不支持三元操作符(三目运算符) "a > b ? a : b"。
}


/****************** if...else 语句*********************/
/*
	if 布尔表达式 {
	   //在布尔表达式为 true 时执行
	} else {
		//在布尔表达式为 false 时执行
	}
 */
func main12() {
	x := 99

	if x < 100 {
		fmt.Printf("x 小于 100\n" )
	} else {
		fmt.Printf("x 不小于 100\n" )
	}
	fmt.Printf("x 的值为 : %d\n", x)
	//x 小于 100
	//x 的值为 : 99
}


/******************if 嵌套语句*********************/
/*
	if 布尔表达式 1 {
		//在布尔表达式 1 为 true 时执行
		if 布尔表达式 2 {
			//在布尔表达式 2 为 true 时执行
		}
	}
 */

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
		}
	}
	fmt.Printf("a 值为 : %d\n", a )
	fmt.Printf("b 值为 : %d\n", b )
	//a 的值为 100 ， b 的值为 200
	//a 值为 : 100
	//b 值为 : 200
}