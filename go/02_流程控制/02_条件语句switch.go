package main

import "fmt"

/***************** switch 语句**********************/
/*
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。
Go 编程语言中 switch 语句的语法如下：
	switch var1 {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。
类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
 */
func main21() {
	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50, 60, 70 : grade = "C"
	default:
		grade = "D"
	}

	switch  {
		case grade == "A":
			fmt.Println("优秀！")
		case grade == "B", grade == "C":
			fmt.Println("良好！")
		case grade == "D":
			fmt.Println("及格！")
		case grade == "F":
			fmt.Println("不及格！")
		default:
			fmt.Println("差！")
	}

	fmt.Println(grade)
	//优秀！
	//A
}

/********************Type Switch*******************/
/*
switch x.(type){
	case type:
		statement(s)
	case type:
		statement(s) //你可以定义任意个数的case
	default: 可选
		statement(s)
}
*/
func main22()  {
	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}
	//写法一：
	switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型 :%T\r\n", i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case  bool, string:
			fmt.Printf("x 是 bool 或 string 型")
		default:
			fmt.Printf("未知型")
	}
}

func main() {
	var j = 0
	switch j {
		case 0:
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("default")
	}

	var k = 0
	switch k {
		case 0:
			fmt.Println("fallthrough")
			fallthrough
			/*Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
			而如果switch没有表达式，它会匹配true。
			Go里面switch默认相当于每个case最后带有break，
			匹配成功后不会自动向下执行其他case，而是跳出整个switch,
			但是可以使用fallthrough强制执行后面的case代码。*/
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("default")
	}

	var m = 0
	switch m {
		case 0, 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("default")
	}

	//而如果switch没有表达式，它会匹配true。
	var n = 0
	switch { ////省略条件表达式，可当 if...else if...else
		case n > 0 && n < 10:
			fmt.Println("n > 0 and n < 10")
		case n > 10 && n <20:
			fmt.Println("n > 10 and n < 20")
		default:
			fmt.Println("default")
	}
}


