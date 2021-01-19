package main

import "fmt"

/******************语法*********************/
/*
 Golang for支持三种循环方式，包括类似 while 的语法。
for循环是一个循环控制结构，可以执行指定次数的循环。

	for init; condition; post { }
    for condition { }
    for { }
    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。
    for语句执行过程如下：
    ①先对表达式 init 赋初值；
    ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，
	然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
 */
func main41() {
	s := "ixfosa"
	for i, n := 0, len(s); i < n;i++ {
		fmt.Println(string(s[i]))
	}

	n := len(s)
	for n > 0 {    //// 替代 while (n > 0) {}    替代 for (; n > 0;) {}
		fmt.Println(string(s[n-1]), n)
		n--
	}

	/*for  {     // 替代 while (true) {}    // 替代 for (;;) {}
		fmt.Println(s)
	}*/
}
func length(s string) int {
	println("length")
	return len(s)
}
func main42() {
	//不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。
	s := "ixfosa"
	for i, n := 0, length(s); i < n; i++ {
		fmt.Println(i, s[i])
	}
	//0 105
	//1 120
	//2 102
	//3 111
	//4 115
	//5 97
}

func main43() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 4, 5, 6}
	for a := 0; a < 10; a++ {
		fmt.Println("a = ", a)
	}
	//

	for a < b {
		a++
		fmt.Println("for a < b => a ", a)
	}

	for i, x := range numbers {
		fmt.Printf("numbers[%d] = %d", i, x)
	}
	//numbers[0] = 1numbers[1] = 2numbers[2] = 3numbers[3] = 4numbers[4] = 5numbers[5] = 6

}



/******************循环嵌套*********************/
/*
	for [condition |  ( init; condition; increment ) | Range] {
	   for [condition |  ( init; condition; increment ) | Range] {
		  statement(s)
	   }
	   statement(s)
	}
 */
func main()  {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
			if (i%j == 0) {
				break
			}
		}
		if (j > i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

/******************无限循环*********************/

func mai44n() {
	for true {
		fmt.Println("hello go")
	}

	for {
		fmt.Println("hello go")
	}
}