package main

import (
	"fmt"
)

// 1.关于字符串连接，下面语法正确的是？
	// A. str := ‘abc’ + ‘123’
	// B. str := “abc” + “123”
	// str := ‘123’ + “abc”
	// D. fmt.Sprintf(“abc%d”, 123)

// 参考答案及解析：BD。知识点：字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString()等。

func main() {
	// str1 := 'adc' + '123' // invalid character literal (more than one character)

	str2 := "abc" + "123" // abc123

	// cannot convert "abc" to type rune
	// invalid operation: '\u0000' + "abc" (mismatched types rune and string)
	// str3 := '123' + "abc"

	str4 := fmt.Sprintf("abc%d", 123)

	fmt.Println(str2,str4) // abc123
}
