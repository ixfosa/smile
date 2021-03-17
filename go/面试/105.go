package main

import "fmt"

// 下面代码输出什么？
func main() {
	isMatch := func(n int) bool {
		switch (n) {
		case 1:
		case 2:     // case 1, 2
			return true
		}
		return false
	}

	fmt.Println(isMatch((1))) // false
	fmt.Println(isMatch(2))   // true

}

// Go 语言的 switch 语句虽然没有”break”，但如果 case 完成程序会默认 break，可以在 case 语句后面加上关键字 fallthrough，
// 这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。或者，利用 case 可以匹配多个值的特性。