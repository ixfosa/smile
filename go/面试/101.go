package main

import "fmt"

// 下面的代码输出什么？
	// 参考答案及解析：编译错误。
		// 很多语言都是采用 ~ 作为按位取反运算符，Go 里面采用的是 ^ 。
func main() {
	// fmt.Println(~2) // bitwise complement operator is ^
	fmt.Println(^2) // -3


	// 按位取反之后返回一个每个 bit 位都取反的数，对于有符号的整数来说，是按照补码进行取反操作的, 快速计算方法：
	// 对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反。例如：
	var a int8 = 3
	var b uint8 = 3
	var c int8 = -3

	fmt.Printf("^%b=%b %d\n", a, ^a, ^a) // ^11=-100 -4
	fmt.Printf("^%b=%b %d\n", b, ^b, ^b) // ^11=11111100 252
	fmt.Printf("^%b=%b %d\n", c, ^c, ^c) // ^-11=10 2


	// 如果作为二元运算符，^ 表示按位异或，即：对应位相同为 0，相异为 1。例如：
	var d int8 = 3
	var f int8 = 5
	fmt.Printf("a: %08b\n",d)			// a: 00000011
	fmt.Printf("c: %08b\n",f)			// c: 00000101
	fmt.Printf("a^c: %08b\n",d ^ f)		// a^c: 00000110

	// 操作符 &^，按位置零，
	// 例如：z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。

	// 或操作符| ，表达式 z = x | y，如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，
	// 否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 &^ 完全相反。
	var x uint8 = 214
	var y uint8 = 92
	fmt.Printf("x: %08b\n",x)			// x: 11010110
	fmt.Printf("y: %08b\n",y)			// y: 01011100
	fmt.Printf("x | y: %08b\n",x | y)	// x | y: 11011110
	fmt.Printf("x &^ y: %08b\n",x &^ y)	// x &^ y: 10000010
}
